package main

import (
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// проверяем нет ли ошибок при выполнения команды
func checkExecuteCmd(cmd exec.Cmd) {
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(out))
	} else {
		fmt.Println(string(out))
	}
}

// трассировка ошибки
func printStackError(err error) {
	stack := err.(stackTracer).StackTrace()
	fmt.Println(stack)
}
