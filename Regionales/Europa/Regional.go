package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"

	pb "github.com/Sistemas-Distribuidos-2023-02/Grupo27-Laboratorio-1/protos"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var llaves int
var registrados int
var numero int

type server struct{
    pb.UnimplementedInteresadosServer
}

func (s *server) Registrados(ctx context.Context, req *pb.NumberRequest) (*pb.NumberResponse, error) {
    receivedNumber := req.GetNumber() 
    return &pb.NumberResponse{Result: "Registrados: "+ strconv.Itoa(int(receivedNumber))}, nil
}

func main() {
    
    hostQ := "dist105"                                             
    qName := "testing"
    txt := 0
    
    // Conexion RabbitMQ
    connQ, err := amqp.Dial("amqp://test:test@" + hostQ + ":5672")
    if err != nil {
        log.Fatal(err)
    } 
    defer connQ.Close()

    //Apertura de canal para procesamiento de mensajes.
    ch, err := connQ.Channel() 
    if err != nil {
        log.Fatal(err)
    } 
    defer ch.Close() 

    for {

        //Mensaje sincrono gRPC
        //Central -> Regional
	    lis, err := net.Listen("tcp", ":50051")
	    if err != nil {
	        log.Fatalf("failed to listen: %v", err)
	    }
	    s := grpc.NewServer()
	    pb.RegisterInteresadosServer(s, &server{})
	    if err := s.Serve(lis); err != nil {
	        log.Fatalf("failed to serve: %v", err)
	    }
        //Fin mensaje

    
        //Usuarios interesados
        if txt == 0 {
            content, err := os.ReadFile("parametros_de_inicio.txt")
            if err != nil {
                log.Fatal(err)
            }
            numero, _= strconv.Atoi(string(content))
            txt = 1
        }
        
        num_2 := float64(numero) /2
        p := num_2*0.20
        random := rand.Intn(int((num_2+p)-(num_2-p) + (num_2 - p)))


        //Mensaje asincrono
        //Regional -> RabbitMQ -> Central
        err = ch.Publish("", qName, false, false,
            amqp.Publishing{
                Headers:     nil,
                ContentType: "text/plain",
                Body:        []byte(strconv.Itoa(random)),
                })
        if err != nil {
            log.Fatal(err)
        }
        // Fin de mensaje


        //Mensaje sincrono gRPC
        //Central -> Regional

        //Fin mensaje
        llaves += random - registrados
    }
}
