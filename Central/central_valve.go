package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/MetalDanyboy/Lab1/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {

    rand.Seed(time.Now().UnixNano())
    content, err := os.ReadFile("parametros de inicio.txt")
    if err != nil {
        log.Fatal(err)
    }

    lineas := strings.Split(string(content), "\n")
    rangoLlaves := strings.Split(lineas[0], "-")


    min, _ := strconv.Atoi(rangoLlaves[0])
    max, _ := strconv.Atoi(rangoLlaves[1])
    iterations, _ := strconv.Atoi(lineas[1])
    llaves := rand.Intn(max-min+1) + min
    log.Printf("Llaves: %d\n", llaves)
    contador := 0


    if iterations == -1 {
        for {
            //randomNumber := rand.Intn(max-min+1) + min
            contador++
            fmt.Printf("Generación %d/infinito\n", contador)

            //-------NOTIFICAR A REGIONALES-------
            //Mensaje sincrono gRPC
            addr:="dist105:50051"
            conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
            if err != nil {
                log.Fatalf("did not connect: %v", err)
            }
            //defer conn.Close()
            c := pb.NewInteresadosClient(conn)

            // Contact the server and print out its response.
            ctx, cancel := context.WithTimeout(context.Background(), time.Second)
            //defer cancel()
            r, err := c.Registrados(ctx, &pb.NumberRequest{Notification: "I have Keys ..."})
            if err != nil {
                log.Fatalf("could not greet: %v", err)
            }
            log.Printf("Verification: %s", r.GetResult())

            conn.Close()
            cancel()
            //-----------------------------
        }
        
    } else {
        for i := 0; i < iterations; i++ {
            //randomNumber := rand.Intn(max-min+1) + min
            contador++
            fmt.Printf("Generación %d/%d\n", contador, iterations)

            //Mensaje sincrono gRPC
            addr:="dist105:50051"
            conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
            if err != nil {
                log.Fatalf("did not connect: %v", err)
            }
            defer conn.Close()
            c := pb.NewInteresadosClient(conn)

            // Contact the server and print out its response.
            ctx, cancel := context.WithTimeout(context.Background(), time.Second)
            defer cancel()
            r, err := c.Registrados(ctx, &pb.NumberRequest{Notification: "I have Keys ..."})
            if err != nil {
                log.Fatalf("could not greet: %v", err)
            }
            log.Printf("Verification: %s", r.GetResult())

        }
    } 
    
}


 



