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
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
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

	conn, err:= grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("Connection: ",conn)
}	