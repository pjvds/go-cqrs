package mongodb

import (
	//"bytes"
	//"encoding/json"
	//"errors"
	//"fmt"
	. "github.com/dominikmayer/go-cqrs/storage"
	"github.com/dominikmayer/go-cqrs/storage/serialization"
	//"net/http"
	//"net/url"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type MongoDB struct {
	baseUrl    string
	database   string
	collection string
	//PageSize   int
	//serializer serialization.Serializer
}

func DialMongoDB(url string, database string, collection string, register *serialization.EventTypeRegister) (*MongoDB, error) {
	return &MongoDB{
		baseUrl:    url,
		database:   database,
		collection: collection,
		//serializer: serialization.NewJsonSerializer(register),
	}, nil
}

//type Event struct {
//	EventId   string         //`json:"eventId"`
//	EventType string         //`json:"eventType"`
//	Data      *storage.Event //`json:"data"`
//}

type MongoEvent struct {
	EventStreamId EventStreamId
	Change *EventStreamChange
}

type MemoryBackend struct {
	changes map[EventStreamId][]*EventStreamChange
}

func (store *MongoDB) WriteStream(change *EventStreamChange) error {
	session, err := mgo.Dial(store.baseUrl)
	Log.Debug("Base-URL: %v", store.baseUrl)
	if err != nil {
		return err
	}
	defer session.Close()

	collection := session.DB(store.database).C(store.collection)

	/* events := change.Events
	data := make([]*MongoEvent, len(events))
	Log.Debug("Creating new stream for %v events", len(events))

	for i := 0; i < len(events); i++ {
		//datapoint := data[i]
		//Log.Debug("Datapoint: %v", datapoint)
		datapoint := &MongoEvent{
			EventStreamId: change.StreamId,
			Change: change,
		}
		Log.Debug("Datapoint: %v", datapoint)
		//datapoint.EventStreamId = change.StreamId
		//datapoint.Change = change
		data = append(data, datapoint)

		e := events[i]

		data[i] = &Event{
			EventId:   e.EventId.String(),
			EventType: e.Name.String(),
			Data:      e,
		}
		Log.Debug("Data %v: %v", i, data[i])
	}*/

	data := change

	Log.Debug("Inserting data: %v", data)
	err = collection.Insert(data)
	if err != nil {
		return err
	}

	return nil
}

func (store *MongoDB) ReadStream(streamId EventStreamId) ([]*Event, error) {
	//events := make([]*Event, 0)
	events := make([]*EventStreamChange,0)
	session, err := mgo.Dial(store.baseUrl)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	collection := session.DB(store.database).C(store.collection)

	Log.Debug("Stream-ID: %v", streamId)
	marshallid, marshallerr := streamId.MarshalJSON()
	Log.Debug("Stream-ID: %v, Error: %v", marshallid, marshallerr)
	err = collection.Find(bson.M{"events.eventid": marshallid}).All(&events)
	//err = collection.Find(bson.M{"events.name": "github.com/dominikmayer/go-cqrs/tests/events/UsernameChanged"}).All(&events)
	Log.Debug("%v Events: %v", len(events), events)
	if err != nil {
		return nil, err
	}
	//events = events.Events
	receivedEvents := events[0].Events
	Log.Debug("%v Events: %v", len(receivedEvents), receivedEvents)
	//err = collection.FindId(streamId).All(&events)

//	for pointer != nil {
//		data, err := pointer.DownloadEvent()
//		if err != nil {
//			return nil, err
//		}
//
//		// Todo: remove eventname, this is not needed any more.
//		event, err := store.serializer.Deserialize(*storage.NewEventName(""), data)
//		if err != nil {
//			return nil, err
//		}
//
//		events = append(events, event)
//		next, err := pointer.Next()
//
//		pointer = next
//	}

	return receivedEvents, nil
}
