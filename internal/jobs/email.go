package jobs

import (

)

type EmailJob struct {

}

func(job EmailJob) Execute(task []byte) error {
	return nil
}