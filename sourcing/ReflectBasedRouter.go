package sourcing

import (
	"reflect"
	"strings"
)

var (
	methodHandlerPrefix = "Handle"
	cache               map[reflect.Type]HandlersMap
)

type HandlersMap map[EventName]func(source interface{}, event Event)

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

	var handlers HandlersMap
	if value, ok := cache[sourceType]; ok {
		handlers = value
	} else {
		// TODO: Now namer could change between entries
		handlers = createEventHandlersForType(namer, sourceType)
	}

	return &ReflectBasedRouter{
		handlers: handlers,
		source:   source,
	}
}

func createEventHandlersForType(namer EventNamer, sourceType reflect.Type) HandlersMap {
	handlers := make(HandlersMap)

	// Loop through all the methods of the source
	for i := 0; i < sourceType.NumMethod(); i++ {
		method := sourceType.Method(i)

		// Only match methods that satisfy prefix
		if strings.HasPrefix(method.Name, methodHandlerPrefix) {
			// Handling methods are defined in code by:
			//   func (source *MySource) HandleMyEvent(e MyEvent).
			// When getting the type of this methods by reflection the signature
			// is as following:
			//   func HandleMyEvent(source *MySource, e MyEvent).
			if method.Type.NumIn() == 2 {
				handler := createEventHandler(method)
				eventName := namer.GetEventNameFromType(method.Type.In(1))
				handlers[eventName] = handler

				Log.Debug("Registered %v as event handler for %v", method.Type.String(), eventName)
			}
		}
	}

	return handlers
}

func createEventHandler(method reflect.Method) func(interface{}, Event) {
	return func(source interface{}, event Event) {
		sourceValue := reflect.ValueOf(source)
		eventValue := reflect.ValueOf(event)

		// Call actual event handling method.
		method.Func.Call([]reflect.Value{sourceValue, eventValue})
	}
}

func (router *ReflectBasedRouter) Route(e EventEnvelope) {
	Log.Debug("Routing %+v", e)

	if handler, ok := router.handlers[e.Name]; ok {
		handler(router.source, e.Payload)
	} else {
		Log.Error("No handler found for event: %v", e.Name)
	}
}
