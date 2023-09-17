package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "github.com/MetalDanyboy/Lab1/protos"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConexionGRPC2(keys int, servidor string) {
	num_cola++
	var puerto, nombre, host string

	if servidor == "America" {
		host = "dist105.inf.santiago.usm.cl"
		puerto = "50052"
		nombre = "America"
	} else if servidor == "Asia" {
		host = "dist106.inf.santiago.usm.cl"
		puerto = "50053"
		nombre = "Asia"
	} else if servidor == "Europa" {
		host = "dist107.inf.santiago.usm.cl"
		puerto = "50054"
		nombre = "Europa"
	} else if servidor == "Oceania" {
		host = "dist108.inf.santiago.usm.cl"
		puerto = "50055"
		nombre = "Oceania"
	}
	log.Println("Connecting to server " + nombre + ": " + host + ":" + puerto + ". . .")
	conn, err := grpc.Dial(host+":"+puerto, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	fmt.Printf("Esperando\n")
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)
	for {
		log.Println("Sending message to server " + nombre + ": " + strconv.Itoa(keys))
		response, err := c.SendKeys(context.Background(), &pb.NumberRequest{Number: int32(keys)})
		if err != nil {
			log.Println("Server " + nombre + " not responding: ")
			log.Println("Trying again in 10 seconds. . .")
			time.Sleep(10 * time.Second)
			continue
		}
		log.Printf("Response from server " + nombre + ": " + "%s", response.Response)
		break
	}
}

func ConexionGRPC(mensaje string, servidor string, wg *sync.WaitGroup) {
	num_cola++
	var puerto, nombre, host string

	if servidor == "America" {
		host = "dist105.inf.santiago.usm.cl"
		puerto = "50052"
		nombre = "America"
	} else if servidor == "Asia" {
		host = "dist106.inf.santiago.usm.cl"
		puerto = "50053"
		nombre = "Asia"
	} else if servidor == "Europa" {
		host = "dist107.inf.santiago.usm.cl"
		puerto = "50054"
		nombre = "Europa"
	} else if servidor == "Oceania" {
		host = "dist108.inf.santiago.usm.cl"
		puerto = "50055"
		nombre = "Oceania"
	}
	log.Println("Connecting to server " + nombre + ": " + host + ":" + puerto + ". . .")
	conn, err := grpc.Dial(host+":"+puerto, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	fmt.Printf("Esperando\n")
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)
	for {
		log.Println("Sending message to server " + nombre + ": " + mensaje)
		response, err := c.SayHello(context.Background(), &pb.Message{Body: mensaje})
		if err != nil {
			log.Println("Server " + nombre + " not responding: ")
			log.Println("Trying again in 10 seconds. . .")
			time.Sleep(10 * time.Second)
			continue
		}
		log.Printf("Response from server " + nombre + ": " + "%s", response.Body)
		break
	}
	defer wg.Done()
}

func esperarHastaCuatro(canal chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		valor := <-canal
		if valor == 4 {
			break
		}
		num_cola++
	}
}

func main() {
	// ... (código anterior sin cambios)

	var llaves int
	for {
		contador++
		if iterations != -1 {
			if contador == iterations+1 {
				break
			}
		}

		if iterations == -1 {
			fmt.Printf("\nGeneración %d/infinito\n", contador)
		} else {
			fmt.Printf("\nGeneración %d/%d\n", contador, iterations)
		}
		llaves = rand.Intn(max-min) + min

		log.Printf("Llaves disponibles: %d\n\n", llaves)

		var wg sync.WaitGroup
		wg.Add(1)
		go ConexionGRPC("LLaves Disponibles", "America", &wg)
		wg.Add(1)
		go ConexionGRPC("LLaves Disponibles", "Asia", &wg)
		wg.Add(1)
		go ConexionGRPC("LLaves Disponibles", "Europa", &wg)
		wg.Add(1)
		go ConexionGRPC("LLaves Disponibles", "Oceania", &wg)
		wg.Wait()

		//Mensaje Rabbit
		forever := make(chan bool)
		var wg2 sync.WaitGroup
		canal := make(chan int)

		wg2.Add(1)
		go esperarHastaCuatro(canal, &wg2)
		go func() {
			for msg := range msgs {
				fmt.Printf("Received Message: %s\n", msg.Body)
				subcadenas := strings.Split(string(msg.Body), "-")

				llaves_pedidas, _ := strconv.Atoi(subcadenas[1])
				if llaves_pedidas > llaves {
					llaves_pedidas = llaves
				}
				if llaves != 0 {
					llaves -= llaves_pedidas
				}

				fmt.Printf("Mensaje asíncrono de servidor %s leído\n", subcadenas[0])
				ConexionGRPC2(llaves_pedidas, subcadenas[0])

				forever <- true
				fmt.Printf("Se inscribieron %d cupos de servidor %s\n", llaves_pedidas, subcadenas[0])
			}
		}()
		<-forever

		if num_cola >= 4 {
			canal <- num_cola
			num_cola = 0
		}
		close(canal)
		wg2.Wait()
	}
	defer log.Println("Closing Central. . .\n")
}

