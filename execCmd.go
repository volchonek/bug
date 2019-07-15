package main

import (
	"os/exec"
	"runtime"
	"sync"
)

// задаем количество одновременно запущенных процессов
const gorutinesNum = 1

// worker для выполнения команды, где chCmd команда для исполнения
func workerCmd(chCmd chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	cmdExec := <-chCmd
	cmdStruct := exec.Command("sh", "-c", cmdExec)
	checkExecuteCmd(*cmdStruct)
	runtime.Gosched()
}

// исполненние команды переданной через cmd
func executeCmd(cmd string) {

	chCmd := make(chan string, 1)
	chCmd <- cmd
	wg := &sync.WaitGroup{}
	for i := 0; i < gorutinesNum; i++ {
		wg.Add(1)
		go workerCmd(chCmd, wg)
	}
	close(chCmd)
}
