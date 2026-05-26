package queue

import (
	"gopeg/preset"
)

type Job struct {
	inputPath  string
	outputPath string
	preset     *preset.Preset
}

type Queue struct {
	IsLocked bool
	jobs     []Job
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) AddJob(inputPath string, preset *preset.Preset) *Queue {
	newJob := Job{
		inputPath:  inputPath,
		outputPath: preset.OutputPath(inputPath),
		preset:     preset,
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

		for _, job := range q.jobs {
			args := job.preset.Args(job.inputPath, job.outputPath)

			// for now, we always overwrite fuck it
			if job.preset.Binary.OverwriteFlag != "" {
				args = append(
					[]string{
						job.preset.Binary.OverwriteFlag,
					},
					args...,
				)
			}

			run(job.preset, args)
		}

		q.Purge()
	}()
}
