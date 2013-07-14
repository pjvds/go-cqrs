package main

import (
	"fmt"
	"github.com/pjvds/gcqrs/eventing"
	"github.com/pjvds/gcqrs/example/domain"
)

func main() {
	user := domain.NewUser("pjvds")
	state := eventing.DefaultContext.GetState(user)

	for _, e := range state.Events {
		fmt.Printf("Event: %s\n", e)
	}

	println("Bye!!")
}
