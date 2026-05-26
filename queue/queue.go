package queue

import (
	"gopeg/preset"
)

type Job struct {
	InputPath  string
	OutputPath string
	Preset     *preset.Preset

	IsRunning bool
	IsDone    bool
}

type Queue struct {
	IsLocked bool
	jobs     []Job
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Length() int {
	return len(q.jobs)
}

func (q *Queue) Job(index int) Job {
	return q.jobs[index]
}

func (q *Queue) AddJob(inputPath string, preset *preset.Preset) *Queue {
	newJob := Job{
		InputPath:  inputPath,
		OutputPath: preset.OutputPath(inputPath),
		Preset:     preset,
	}
	q.jobs = append(q.jobs, newJob)
	return q
}

func (q *Queue) Purge() *Queue {
	q.jobs = nil
	return q
}

func (q *Queue) lock() *Queue {
	q.IsLocked = true
	return q
}

func (q *Queue) unlock() *Queue {
	q.IsLocked = false
	return q
}

func (q *Queue) Execute() {
	if q.IsLocked {
		return
	}

	// all jobs inside a thread
	go func() {
		q.lock()
		defer q.unlock()

		for i := range q.jobs {
			args := q.jobs[i].Preset.Args(q.jobs[i].InputPath, q.jobs[i].OutputPath)

			// for now, we always overwrite fuck it
			if q.jobs[i].Preset.Binary.OverwriteFlag != "" {
				args = append(
					[]string{
						q.jobs[i].Preset.Binary.OverwriteFlag,
					},
					args...,
				)
			}

			q.jobs[i].IsRunning = true
			run(q.jobs[i].Preset, args)
			q.jobs[i].IsRunning = false
			q.jobs[i].IsDone = true
		}

		q.Purge()
	}()
}
