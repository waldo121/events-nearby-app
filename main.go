package main

import (
	"fmt"
	"github.com/waldo121/events-nearby-app/clients"
)

func main() {
	fmt.Println(clients.GetEvents(1))
}
