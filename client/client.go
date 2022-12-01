package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
  rpcClient, err := rpc.Dial("tcp", "localhost:1234")
  if err != nil {
    log.Fatal(err)
  }

  var reply string
  // call remote method: HelloService.Hello(Jack)
  err = rpcClient.Call("HelloService.Hello", "Jack", &reply)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(reply)
}
