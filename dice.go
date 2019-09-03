package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 6
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(n) + 1)
}
