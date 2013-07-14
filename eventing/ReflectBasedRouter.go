package eventing

import (
	"reflect"
	"strings"
)

var MethodHandlerPrefix = "Handle"

// Routes events to methods of an struct by convention. There should be one
// router per event source instance.
//
// The convention is: HandleXXX(e EventType)
type ReflectBasedRouter struct {
	handlers map[EventName]EventHandler
	source   interface{}
}

func NewReflectBasedRouter(namer EventNamer, source interface{}) EventRouter {
	sourceType := reflect.TypeOf(source)
	handlers := make(map[EventName]EventHandler)

	for i := 0; i < sourceType.NumMethod(); i++ {
		method := sourceType.Method(i)

		// Handler method will have the following signature from an reflection
		// point of view: HandleUserCreated(*domain.User, events.UserCreated)
		if strings.HasPrefix(method.Name, MethodHandlerPrefix) && method.Type.NumIn() == 2 {
			handler := func(e Event) {
				method.Func.Call([]reflect.Value{reflect.ValueOf(source), reflect.ValueOf(e)})
			}
			eventName := namer.GetEventNameFromType(method.Type.In(1))
			handlers[eventName] = handler
		}
	}

	return &ReflectBasedRouter{
		handlers: handlers,
	}
}

func (router *ReflectBasedRouter) Route(e EventEnvelope) {
	Log.Debug("Routing %+v", e)

	if handler, ok := router.handlers[e.Name]; ok {
		handler(e.Payload)
	} else {
		Log.Error("No handler found for event: %v", e.Name)
	}
}
