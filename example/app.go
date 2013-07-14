package main

import (
	"github.com/pjvds/gcqrs/eventing"
	"github.com/pjvds/gcqrs/example/domain"
)

func main() {
	InitLogging()

	user := domain.NewUser("pjvds")
	state := eventing.DefaultContext.GetState(user)

	for index, value := range state.Events {
		Log.Info("Event %v: %v\n", index+1, value)
	}

	Log.Notice("Bye!!")
}
