package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MetalDanyboy/Lab1/protos"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {

    //Conexion Grpc
	//##############################################
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	//##############################################


	//Enviar Mensaje 1 - Grpc (Tengo llaves)
	//##############################################
	response, err := c.SayHello(context.Background(), &pb.Message{Body: "Tengo llaves"})
	if err != nil {
		log.Fatalf("Error send msj: %s", err)
	}
	log.Printf("Response from regional: %s", response.Body)
	//##############################################


    //Conexion Rabbit (en la misma VM)
    //##############################################
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		"testing", // queue
		"",        // consumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       //args
	)
	if err != nil {
		panic(err)
	}
	//##############################################


	//Recibir Mensaje  - Rabbit (Usuarios interesados)
	//##############################################
	forever := make(chan bool)
	mensaje := make(chan string)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
			mensaje <- string(msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
	//##############################################


	//Enviar Mensaje 2 - Grpc (n° de registrados)
	//##############################################
	mensaje2:= <- mensaje
	response, err = c.SayHello(context.Background(), &pb.Message{Body: mensaje2 +" Msg 2"})
	if err != nil {
		log.Fatalf("Error send msj: %s", err)
	}
	log.Printf("Response from regional: %s", response.Body)
}