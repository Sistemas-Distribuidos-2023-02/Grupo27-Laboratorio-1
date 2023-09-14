package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"

	pb "github.com/MetalDanyboy/Lab1/protos"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

//var llaves int

type Server struct {
	pb.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}


func main() {

    //Grpc
	//##############################################
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//s := chat.Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	//##############################################



	//Usuarios
	usuarios := rand.Intn(100)


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

	// declaring queue with its properties over the the channel opened
	queue, err := channel.QueueDeclare(
		"testing", // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		panic(err)
	}
    //##############################################


	//Enviar Rabbit (Usuarios interesados)
	//##############################################
	err = channel.Publish(
		"",        // exchange
		"testing", // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(strconv.Itoa(usuarios)),
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)



    //Grpc
	//##############################################
	lis, err = net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//s := chat.Server{}

	grpcServer = grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
	//##############################################