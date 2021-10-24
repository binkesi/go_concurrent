package main

import (
	"GoConcurrent/concurrent_models"
	"fmt"
)

func main() {
	go concurrent_models.Setup()
	for !concurrent_models.Done {
	}
	fmt.Println(concurrent_models.A)
}
