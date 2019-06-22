package main

import (
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// проверяем ошибку после выполнения команды и выоводим результат
func checkExecuteCmd(cmd exec.Cmd) {
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(out))
		printStackError(err)
	} else {
		fmt.Println(string(out))
	}
}

// получаем трассировку ошибки err и выводим на экран
func printStackError(err error) {
	stack := err.(stackTracer).StackTrace()
	fmt.Println(stack)
}
