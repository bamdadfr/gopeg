package queue

import (
	"context"
	"gopeg/pipeline"
	"gopeg/preset"
	"log"
	"os"
	"sync"
)

type Job struct {
	InputPath  string
	OutputPath string
	Preset     *preset.Preset

	IsRunning bool
	IsDone    bool
	HasError  bool

	// Pipeline tracking: jobs sharing the same non-zero PipelineID
	// are dependent — if one fails, the rest are skipped.
	PipelineID    int
	Intermediates []string // paths to remove after the group finishes
}

type Queue struct {
	mu             sync.Mutex
	isLocked       bool
	jobs           []Job
	onUpdate       func()
	ctx            context.Context
	nextPipelineID int
}

func NewQueue(ctx context.Context) *Queue {
	return &Queue{ctx: ctx}
}

func (q *Queue) OnNotify(f func()) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.onUpdate = f
}

func (q *Queue) notify() {
	q.mu.Lock()
	f := q.onUpdate
	q.mu.Unlock()
	if f != nil {
		f()
	}
}

func (q *Queue) Length() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.jobs)
}

func (q *Queue) Job(index int) Job {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.jobs[index]
}

func (q *Queue) IsRunning() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.isLocked
}

func (q *Queue) AddJob(inputPath string, p *preset.Preset) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.isLocked {
		return
	}

	q.jobs = append(q.jobs, Job{
		InputPath:  inputPath,
		OutputPath: p.OutputPath(inputPath),
		Preset:     p,
	})
}

func (q *Queue) AddPipeline(inputPath string, p *pipeline.Pipeline) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.isLocked {
		return
	}

	q.nextPipelineID++
	pid := q.nextPipelineID

	steps := p.Steps
	var intermediates []string
	currentInput := inputPath

	for i := range steps {
		step := &steps[i]
		isLast := i == len(steps)-1

		var outputPath string
		if isLast {
			outputPath = p.FinalPath(inputPath)
		} else {
			outputPath = pipeline.IntermediatePath(inputPath, *step)
			intermediates = append(intermediates, outputPath)
		}

		job := Job{
			InputPath:  currentInput,
			OutputPath: outputPath,
			Preset:     step,
			PipelineID: pid,
		}

		// Last job in the pipeline carries the cleanup list
		if isLast {
			job.Intermediates = intermediates
		}

		q.jobs = append(q.jobs, job)
		currentInput = outputPath
	}
}

func (q *Queue) Purge() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.jobs = nil
}

func (q *Queue) Execute() {
	q.mu.Lock()
	if q.isLocked {
		q.mu.Unlock()
		return
	}
	q.isLocked = true
	q.mu.Unlock()

	go func() {
		defer func() {
			q.mu.Lock()
			q.isLocked = false
			q.mu.Unlock()
		}()

		failedPipelines := map[int]bool{}

		for i := range q.jobs {
			q.mu.Lock()
			job := q.jobs[i]
			q.mu.Unlock()

			// Skip remaining steps of a failed pipeline
			if job.PipelineID != 0 && failedPipelines[job.PipelineID] {
				q.mu.Lock()
				q.jobs[i].IsDone = true
				q.jobs[i].HasError = true
				q.mu.Unlock()
				q.notify()
				continue
			}

			passes := job.Preset.Args(job.InputPath, job.OutputPath)

			q.mu.Lock()
			q.jobs[i].IsRunning = true
			q.mu.Unlock()
			q.notify()

			var jobFailed bool
			for _, passArgs := range passes {
				if job.Preset.Binary.OverwriteFlag != "" {
					passArgs = append([]string{job.Preset.Binary.OverwriteFlag}, passArgs...)
				}

				if err := run(q.ctx, job.Preset, passArgs); err != nil {
					log.Println("Aborting job:", err)
					jobFailed = true
					break
				}
			}

			q.mu.Lock()
			q.jobs[i].IsRunning = false
			q.jobs[i].IsDone = true
			q.jobs[i].HasError = jobFailed
			q.mu.Unlock()
			q.notify()

			if jobFailed && job.PipelineID != 0 {
				failedPipelines[job.PipelineID] = true
			}

			// Clean up intermediate files after last pipeline step
			if len(job.Intermediates) > 0 && !jobFailed {
				for _, path := range job.Intermediates {
					if err := os.Remove(path); err != nil {
						log.Println("Cleanup:", err)
					}
				}
			}
		}

		q.Purge()
	}()
}
