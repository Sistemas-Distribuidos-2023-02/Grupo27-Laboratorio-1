package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
	"os"
	"os"
	"strconv"
	"time"

	//pb "Tarea1-Grupo27/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	//"google.golang.org/grpc"
)


func main() {
	
	hostQ := "dist105"                                             
	qName := "testing"
	numero := -1
	
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

		//Fin mensaje


		//Usuarios interesados
		if numero == -1 {
    		content, err := os.ReadFile("parametros_de_inicio.txt")
    		if err != nil {
        		log.Fatal(err)
    		}
    		num, _ := strconv.Atoi(content)
    		min := int((num/2)*0.80)
    		max := int((num/2)*1.20)
    		numero ++
    	}
    	numero += rand.Intn(max-min+1) + min


		//Mensaje asincrono
		//Regional -> RabbitMQ -> Central
		err = ch.Publish("", qName, false, false,
			amqp.Publishing{
				Headers:     nil,
				ContentType: "text/plain",
				Body:        []byte(strconv.Itoa(numero)),
				})
		if err != nil {
			log.Fatal(err)
		}
		// Fin de mensaje


		// notificacion desde la central que le indique cuantos usuarios 
		// no pudieron ser registrados, para poder restarlos respecto al 
		// valor enviado
		
		//numero -= 
	}
}
