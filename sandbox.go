package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// // печать аргументов
// func argsPrint(args []string) {
// 	fmt.Printf("All arguments command line: %v\n", args)
// 	fmt.Printf("Worker arguments command line: %v\n", args[1:])
// }

// // печать аргументов идущих за флагом
// func flagsPrint(args []string) {
// 	flag.String("test", args[1], "a string type")
// 	flag.String("test1", args[2], "a string type")
// 	flag.Parse()
// 	fmt.Printf("Flags arguments command line: %v\n", flag.Args())
// }

// печать перменных окружения
func enviromentsPrint() {
	fmt.Printf("Enviroments:\n")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Printf("Enviroment: %v\n", pair)
	}
}

func executeCmdRun() {
	// _, err := exec.Command("sh", "-c", "sudo docker build . -t golang:bug").Output()
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// }

	_, err := exec.Command("sh", "-c", "sudo docker run -it --name=golang-bug golang:bug").Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	os.Setenv("flag", "true")
}

func sandBox() {
	// // чтение и печать аргументов из командной строки
	// argsAll := os.Args
	// argsPr := argsPrint
	// argsPr(argsAll)

	// // чтение и печать аргументов из команндной строки стоящих зафлагом
	// flagPr := flagsPrint
	// flagPr(argsAll)

	// установка и печать переменных окружения
	// os.Setenv("buildEnv", "sudo docker build . -t golang:testSr1")
	// os.Setenv("runEnv", "sudo docker run -it --name=golangSr1 golang:testSr1")
	// enviromentsPrint := enviromentsPrint
	// enviromentsPrint()

	// запуск на сборку билда

	// os.Setenv("path", ".")
	// os.Setenv("flag", "-t")
	// os.Setenv("box", "golang:testSr1")

	if os.Getenv("flag") == "flag" {
		fmt.Printf("It is run, well done flag=%v", os.Getenv("flag"))
	}

	fmt.Printf("test")

	executeCmdRun()
}
