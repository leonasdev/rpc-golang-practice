package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, replay *string) error {
  *replay = "hello: " + request
  return nil
}

func main() {
  rpc.Register(new(HelloService))

  tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
  if err != nil {
    log.Fatal(err)
  }

  listner, err := net.ListenTCP("tcp", tcpAddr)
  if err != nil {
    log.Fatal(err)
  }

  for {
    conn, err := listner.Accept()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println("Recive request from: ", conn.RemoteAddr().String())

    go rpc.ServeConn(conn)
  }
}
