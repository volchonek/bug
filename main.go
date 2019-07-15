package main

import (
	"fmt"
	"os"
)

const (
	grpc = "grpc"
	rest = "rest"

	cmdBuild = "sudo docker build ~/go_projects/src/btest -t golang:btest"
	// to do:	cmdPull = "sudo docker pull ..."
	cmdRun = "sudo docker run --rm -i --name=btest -p 8082:80 golang:btest"
)

func main() {
	// sandBox()

	// собираем образ
	executeCmd(cmdBuild)
	// запускаем образ
	executeCmd(cmdRun)

	fmt.Println("Desctop is run!!!")

	if os.Args[1] == grpc {
		// runGrpcServer()
	}

	if os.Args[1] == rest {
		runHttpServer(":8081")
	}
}
