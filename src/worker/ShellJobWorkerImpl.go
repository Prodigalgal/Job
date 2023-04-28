package worker

import (
	"Job/src/beans"
	"os"
)

type ShellJobWorker struct {
	Job    *beans.Job
	File   *os.File
	Result map[string]interface{}
}

func (s ShellJobWorker) Execute() {
	//	执行脚本
}
