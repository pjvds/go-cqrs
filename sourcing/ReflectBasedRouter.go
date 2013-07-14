package sourcing

import (
	"reflect"
	"strings"
)

var MethodHandlerPrefix = "Handle"

type HandlersMap map[EventName]EventHandler

// Routes events to methods of an struct by convention. There should be one
// router per event source instance.
//
// The convention is: func(s MySource) HandleXXX(e EventType)
type ReflectBasedRouter struct {
	handlers HandlersMap
	source   interface{}
}

func NewReflectBasedRouter(namer EventNamer, source interface{}) EventRouter {
	sourceType := reflect.TypeOf(source)
	handlers := make(HandlersMap)

	// Loop through all the methods of the source
	for i := 0; i < sourceType.NumMethod(); i++ {
		method := sourceType.Method(i)

		// Only match methods that satisfy prefix
		if strings.HasPrefix(method.Name, MethodHandlerPrefix) {
			// Handling methods are defined in code by:
			//   func (source *MySource) HandleMyEvent(e MyEvent).
			// When getting the type of this methods by reflection the signature
			// is as following:
			//   func HandleMyEvent(source *MySource, e MyEvent).
			if method.Type.NumIn() == 2 {
				handler := createEventHandler(source, method)
				eventName := namer.GetEventNameFromType(method.Type.In(1))
				handlers[eventName] = handler

				Log.Debug("Registered %v as event handler for %v", method.Type.String(), eventName)
			}
		}
	}

	return &ReflectBasedRouter{
		handlers: handlers,
	}
}

func createEventHandler(source interface{}, method reflect.Method) EventHandler {
	return func(event Event) {
		sourceValue := reflect.ValueOf(source)
		eventValue := reflect.ValueOf(event)

		// Call actual event handling method.
		method.Func.Call([]reflect.Value{sourceValue, eventValue})
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
