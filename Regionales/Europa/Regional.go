package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/MetalDanyboy/Lab1/protos"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var llaves int
var registrados int
var numero int
var name string


type server struct{
    pb.UnimplementedInteresadosServer
}

func (s *server) Registrados(ctx context.Context, req *pb.NumberRequest) (*pb.NumberResponse, error) {
    request := req.GetNotification() 
    log.Printf("Request: %s", request)
    /*if request == "I have Keys ..." {
        keys_available <- true
    }*/
    return &pb.NumberResponse{Result: name+": OK"}, nil
}

func main() {
    keys_available := make(chan bool)
    keys_available <- false
    name = "Europa"
    
    rand.Seed(time.Now().UnixNano())
    hostQ := "dist105"                                             
    qName := "testing"
    txt := 0
    
    // Conexion RabbitMQ
    connQ, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672")
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
    
    go func() {
        //Mensaje sincrono gRPC
        //Central -> Regional
        log.Println("Escuchando en puerto 50051 . . .")
        lis, err := net.Listen("tcp", ":50051")
        if err != nil {
            log.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        pb.RegisterInteresadosServer(s, &server{})
        if err := s.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
        keys_available<-true
    }()
       
        //Fin mensaje
    for {
        valorRecibido:= <-keys_available
        if valorRecibido{
            log.Println("Llaves disponibles")
            //Usuarios interesados
            if txt == 0 {
                content, err := os.ReadFile("parametros de inicio.txt")
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
        keys_available<-false
    }
}
