package main

import (
	//"tim"
	//"strconv"
	//"strings"
	//"math"
	//"net"
	//"context"
	//"fmt"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
	//"os/signal"
	//"sync"
)

func main() {
	content, err := os.ReadFile("parametros de inicio.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
	producer()
}

func producer() {
	//TODO
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5000/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(connection)
	fmt.Println("Producer")
}

/*func consumer() {
	//TODO
}*/

