package queue

import (
	"context"
	"gopeg/preset"
	"sync"
)

type Job struct {
	InputPath  string
	OutputPath string
	Preset     *preset.Preset

	IsRunning bool
	IsDone    bool
}

type Queue struct {
	mu       sync.Mutex
	isLocked bool
	jobs     []Job
	onUpdate func()
	ctx      context.Context
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

func (q *Queue) AddJob(inputPath string, preset *preset.Preset) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.isLocked {
		return
	}

	q.jobs = append(q.jobs, Job{
		InputPath:  inputPath,
		OutputPath: preset.OutputPath(inputPath),
		Preset:     preset,
	})
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

		for i := range q.jobs {
			q.mu.Lock()
			job := q.jobs[i]
			q.mu.Unlock()

			passes := job.Preset.Args(job.InputPath, job.OutputPath)

			q.mu.Lock()
			q.jobs[i].IsRunning = true
			q.mu.Unlock()
			q.notify()

			for _, passArgs := range passes {
				if job.Preset.Binary.OverwriteFlag != "" {
					passArgs = append([]string{job.Preset.Binary.OverwriteFlag}, passArgs...)
				}

				run(q.ctx, job.Preset, passArgs)
			}

			q.mu.Lock()
			q.jobs[i].IsRunning = false
			q.jobs[i].IsDone = true
			q.mu.Unlock()
			q.notify()
		}

		q.Purge()
	}()
}
