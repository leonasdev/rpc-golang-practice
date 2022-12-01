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
  // A server registers an object, making it visible as a service with the name of the type of the object.
  rpc.Register(new(HelloService))

  // defined a tcp address
  tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
  if err != nil {
    log.Fatal(err)
  }

  listner, err := net.ListenTCP("tcp", tcpAddr)
  if err != nil {
    log.Fatal(err)
  }

  // start listening
  for {
    conn, err := listner.Accept()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println("Recive request from: ", conn.RemoteAddr().String())

    go rpc.ServeConn(conn)
  }
}
