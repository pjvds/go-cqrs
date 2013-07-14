package eventing

import (
	"reflect"
	"strings"
)

var MethodHandlerPrefix = "Handle"

type ReflectBasedRouter struct {
	handlers map[EventName]EventHandler
	source   interface{}
}

func NewReflectBasedRouter(namer EventNamer, source interface{}) EventRouter {
	t := reflect.TypeOf(source)
	handlers := make(map[EventName]EventHandler)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		if strings.HasPrefix(m.Name, MethodHandlerPrefix) && m.Type.NumIn() == 2 {
			source := source
			handler := func(e Event) {
				m.Func.Call([]reflect.Value{reflect.ValueOf(source), reflect.ValueOf(e)})
			}
			name := namer.GetEventNameFromType(m.Type.In(1))
			handlers[name] = handler

		}
	}

	return &ReflectBasedRouter{
		handlers: handlers,
	}
}

func (router *ReflectBasedRouter) Route(e EventEnvelope) {
	Log.Debug("Routing %+v", e)

	handler := router.handlers[e.Name]
	if handler == nil {
		Log.Error("No handler found for event: %v", e.Name)
	} else {
		handler(e.Payload)
	}
}
