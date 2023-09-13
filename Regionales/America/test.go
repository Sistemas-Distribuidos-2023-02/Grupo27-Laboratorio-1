package main

import (
	//"tim"
	"strconv"
	"strings"

	//"math"
	//"net"
	//"context"
	"fmt"
	"log"
	"os"

	//"os/signal"
	//"sync"
	"google.golang.org/grpc"
)

func main2() {
    content, err := os.ReadFile("parametros de inicio.txt")
    if err != nil {
        log.Fatal(err)
    }

	var initial_users int
	//print line from content
	for _, line := range strings.Split(string(content), "\n") {
		initial_users, err = strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error reading file",err)
			return
		}
	}
	fmt.Println("Initial users: ",initial_users)

	conn, _:= grpc.Dial("localhost:50051", grpc.WithInsecure())
	fmt.Println("Connection: ",conn)
}
