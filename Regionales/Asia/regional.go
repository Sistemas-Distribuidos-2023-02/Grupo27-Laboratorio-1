package main

import (
	"strconv"
	"strings"

	//"math"
	"context"
	"net"

	//"fmt"
	"fmt"
	"log"
	"os"

	//"google.golang.org/grpc"
	//"os/signal"
	//"sync"
	//"google.golang.org/grpc"

	pb "github.com/Sistemas-Distribuidos-2023-02/Grupo27-Laboratorio-1/protos"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
	initialUsers int
}

// SayHello implements protos.GreeterServer
func (s *server) UserRegion(ctx context.Context, in *pb.UserRegionRequest) (*pb.UserRegionReply, error) {
	
	/*// Establece una semilla aleatoria basada en la hora actual.
    rand.Seed(time.Now().UnixNano())

    // Define los valores mínimo y máximo para el rango de números aleatorios.
    min := int(float64(s.initialUsers /2) * 0.8 )
    max := int(float64(s.initialUsers /2) * 1.2)

    // Genera un número aleatorio entre min (incluido) y max (excluido).
    numeroAleatorio := int32(rand.Intn(max-min) + min)
	fmt.Println("Numero Aleatorio: "+ strconv.Itoa(int(numeroAleatorio)))*/
	return &pb.UserRegionReply{Verificacion: "OK"}, nil
}

/* //SayHelloAgain implements helloworld.GreeterServer
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}
*/



func main() {
    content, err := os.ReadFile("parametros de inicio.txt")
    if err != nil {
        log.Fatal(err)
    }
	//fmt.Println(content)


	var initial_users int
	//print line from content
	for _, line := range strings.Split(string(content), "\n") {
		initial_users, err = strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting to Int",err)
			return
		}
	}
	fmt.Println("Initial users: ",initial_users)

	port:= 50051

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}else{
		fmt.Println("Listening on port: ",port)
	}

	grpcServer := grpc.NewServer()
	myServer := &server{initialUsers: initial_users}
	pb.RegisterGreeterServer(grpcServer, myServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}