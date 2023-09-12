package main

import (
    "time"
    "strconv"
    "strings"
    "math"
    "math/rand"
    "net"
    "context"
    "fmt"
    "log"
    "os"
    "os/signal"
    "sync"
)


func main() {

    rand.Seed(time.Now().UnixNano())
    content, err := os.ReadFile("parametros_de_inicio.txt")
    if err != nil {
        log.Fatal(err)
    }

    lineas := strings.Split(string(content), "\n")
    rangoLlaves := strings.Split(lineas[0], "-")


    min, _ := strconv.Atoi(rangoLlaves[0])
    max, _ := strconv.Atoi(rangoLlaves[1])
    iterations, _ := strconv.Atoi(lineas[1])
    llaves := rand.Intn(max-min+1) + min
    contador := 0


    if iterations == -1 {
        for {
            randomNumber := rand.Intn(max-min+1) + min
            contador++
            fmt.Printf("Generación %d/infinito\n", contador)
        }
        
    } else {
        for i := 0; i < iterations; i++ {
            randomNumber := rand.Intn(max-min+1) + min
            contador++
            fmt.Printf("Generación %d/%d"\n, contador, iterations)
        }
    }   
}


 



