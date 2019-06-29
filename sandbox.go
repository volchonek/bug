package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
)

const (
	gorutinesNum = 1
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

// // печать перменных окружения
// func enviromentsPrint() {
// 	fmt.Printf("Enviroments:\n")
// 	for _, e := range os.Environ() {
// 		pair := strings.Split(e, "=")
// 		fmt.Printf("Enviroment: %v\n", pair)
// }

// worker для выполнения команды, где cmd команда для исполнения
func worker(cmd chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	cmdExec := <-cmd
	cmdrun := exec.Command("sh", "-c", cmdExec)
	checkExecuteCmd(*cmdrun)
	runtime.Gosched()
}

// запуск контейнера btest
func executeCmd() {
	// собираем образ
	// cmdbuild := exec.Command("sh", "-c", "sudo docker build ~/go_projects/src/btest -t golang:btest")
	// checkExecuteCmd(*cmdbuild)

	// параллельно разворачиваем образ(ы) в отдельной(ых) горутине(ах), избегаем блокировку в основном потоке
	cmd := make(chan string, 1)
	cmd <- "sudo docker run --rm -i --name=btest -p 8082:80 golang:btest"
	wg := &sync.WaitGroup{}
	for i := 0; i < gorutinesNum; i++ {
		wg.Add(1)
		go worker(cmd, wg)
	}
	close(cmd)
}

func sandBox() {
	// // чтение и печать аргументов из командной строки

	// argsAll := os.Args
	// argsPr := argsPrint
	// argsPr(argsAll)

	// // чтение и печать аргументов из команндной строки стоящих зафлагом
	// flagPr := flagsPrint
	// flagPr(argsAll)

	// // установка и печать переменных окружения

	// os.Setenv("buildEnv", "sudo docker build . -t golang:testSr1")
	// os.Setenv("runEnv", "sudo docker run -it --name=golangSr1 golang:testSr1")
	// enviromentsPrint := enviromentsPrint
	// enviromentsPrint()

	// запуск контейра с тестами
	executeCmd()

	fmt.Println("Desctop is run!!!")

	// запускаем сервис для получения информации по контейнерам
	runServer(":8081")
}
