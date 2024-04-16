package pkg

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	_ "github.com/go-co-op/gocron/v2"
)

type ExecuteJobFn func()

func Job(fn ExecuteJobFn, bParams BackupParams) (gocron.Scheduler, gocron.Job, error) {
	s, _ := gocron.NewScheduler()

	job, err := s.NewJob(
		gocron.CronJob(
			*bParams.crontab, // funciona por causa do WITHSECONDS do "gocron"
			false,
		),
		gocron.NewTask(
			fn,
		),
	)

	// each job has a unique id
	fmt.Println(job.ID())

	// start the scheduler
	s.Start()

	return s, job, err

}
