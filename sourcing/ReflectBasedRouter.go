package sourcing

import (
	"reflect"
	"strings"
)

var (
	methodHandlerPrefix = "Handle"
	cache               map[reflect.Type]handlersMap
)

type handlersMap map[reflect.Type]func(source interface{}, event Event)

// Routes events to methods of an struct by convention. There should be one
// router per event source instance.
//
// The convention is: func(s MySource) HandleXXX(e EventType)
type ReflectBasedRouter struct {
	handlers   handlersMap
	source     interface{}
	sourceType reflect.Type
}

func NewReflectBasedRouter(source interface{}) EventRouter {
	sourceType := reflect.TypeOf(source)

	var handlers handlersMap
	if value, ok := cache[sourceType]; ok {
		handlers = value
	} else {
		// TODO: Now namer could change between entries
		handlers = createEventHandlersForType(sourceType)
	}

	return &ReflectBasedRouter{
		handlers:   handlers,
		source:     source,
		sourceType: sourceType,
	}
}

func createEventHandlersForType(sourceType reflect.Type) handlersMap {
	handlers := make(handlersMap)

	// Loop through all the methods of the source
	methodCount := sourceType.NumMethod()
	for i := 0; i < methodCount; i++ {
		method := sourceType.Method(i)

		// Only match methods that satisfy prefix
		if strings.HasPrefix(method.Name, methodHandlerPrefix) {
			// Handling methods are defined in code by:
			//   func (source *MySource) HandleMyEvent(e MyEvent).
			// When getting the type of this methods by reflection the signature
			// is as following:
			//   func HandleMyEvent(source *MySource, e MyEvent).
			if method.Type.NumIn() == 2 {
				eventType := method.Type.In(1)
				handler := createEventHandler(method)
				handlers[eventType] = handler
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

func (router *ReflectBasedRouter) Route(e Event) {
	Log.Debug("Routing %+v", e)

	eventType := reflect.TypeOf(e)
	if handler, ok := router.handlers[eventType]; ok {
		handler(router.source, e)
	} else {
		Log.Error("No handler found for event: %v in %v", eventType.String(), router.sourceType.String())
	}
}
