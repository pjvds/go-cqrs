package main

import (
	"github.com/pjvds/go-cqrs/eventing"
	"github.com/pjvds/go-cqrs/example/domain"
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
