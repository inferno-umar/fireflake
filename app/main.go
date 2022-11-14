package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"uid_gen/grpc/uid_gen"
	"uid_gen/internal/pkg/global"
	"uid_gen/internal/pkg/grpcServer"

	"google.golang.org/grpc"
)

func main() {
	global.ParseArgs(os.Args)

	go func() {
		reader := bufio.NewScanner(os.Stdin)
		for reader.Scan() {
			switch reader.Text() {
			case "close":
				os.Exit(0)
				break
			}
		}
	}()

	lisG, err := net.Listen("tcp", global.IP+global.Port)
	global.HandleErr(err)

	s := grpc.NewServer()
	uid_gen.RegisterUidGenServer(s, &grpcServer.GRPC_UID_server{})

	log.Println("Listening GRPC ID generator on ", lisG.Addr())
	if err = s.Serve(lisG); err != nil {
		log.Println(err)
	}
}
