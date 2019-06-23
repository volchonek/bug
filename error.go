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
		printError(out, err)
		printStackError(err)
	} else {
		fmt.Println(string(out))
	}
}

// трассировка ошибки
func printStackError(err error) {
	stack := err.(stackTracer).StackTrace()
	fmt.Println(stack)
}

// содержимое ошибки
func printError(args ...interface{}) {

	var err error
	var out []byte

	if 1 > len(args) {
		fmt.Println("Error")
		return
	}

	for id, arg := range args {
		switch id {
		case 0: // err
			param, ok := arg.(error)
			if !ok {
				fmt.Println(errors.New("1st parametr not type error"))
			}
			err = param
			return
		case 1: // out
			param, ok := arg.([]byte)
			if !ok {
				fmt.Println(errors.New("2st parametr not type []byte"))
			}
			out = param
			return
		default:
			fmt.Println(errors.New("Wrong parameter count"))
			return
		}
	}

	fmt.Println(fmt.Sprint(err) + ": " + string(out))
	return
}
