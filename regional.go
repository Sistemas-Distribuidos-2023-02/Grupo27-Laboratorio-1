package main

import (
	"tim"
	"strconv"
	"strings"
	"math"
	"net"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
    content, err := os.ReadFile("parametros de inicio.txt")
    if err != nil {
        log.Fatal(err)
    }

}