package main

import (
	"net/http"
	"os"
	"log"
	"net"
	"strconv"
	"fmt"
)

func main() {
	argsLen := len(os.Args)
	var arg string
	if argsLen > 1 {
		arg = os.Args[1]
		if _, err := strconv.ParseInt(arg,10,16); err != nil {
			fmt.Println("Port ", arg, "is illegal (should be an uint16)")
			fmt.Printf("Usage: %s [port]", os.Args[0])
			os.Exit(1)
			return
		}
	} else {
		arg = "0" // any
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working dir ", err)
		return
	}

	listener, err := net.Listen("tcp", ":"+arg)
	if err != nil {
		log.Fatal("Failed to listen", err)
		return
	}
	log.Println("Listening on", listener.Addr())
	log.Fatal(http.Serve(listener, http.FileServer(http.Dir(cwd))))
}


