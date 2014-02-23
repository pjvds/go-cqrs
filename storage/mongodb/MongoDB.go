package mongodb

import (
	//"bytes"
	//"encoding/json"
	//"errors"
	//"fmt"
	"github.com/dominikmayer/go-cqrs/storage"
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

type Event struct {
	EventId   string         //`json:"eventId"`
	EventType string         //`json:"eventType"`
	Data      *storage.Event //`json:"data"`
}

func (store *MongoDB) WriteStream(change *storage.EventStreamChange) error {
	session, err := mgo.Dial(store.baseUrl)
	Log.Debug("Base-URL: %v", store.baseUrl)
	if err != nil {
		return err
	}
	defer session.Close()

	collection := session.DB(store.database).C(store.collection)

	events := change.Events
	data := make([]*Event, len(events))
	Log.Debug("Creating new stream for %v events", len(events))

	for i := 0; i < len(events); i++ {
		e := events[i]

		data[i] = &Event{
			EventId:   e.EventId.String(),
			EventType: e.Name.String(),
			Data:      e,
		}
		Log.Debug("Data %v: %v", i, data[i])
	}

	Log.Debug("Inserting data: %v", data)
	err = collection.Insert(events)
	if err != nil {
		return err
	}

	return nil
}

func (store *MongoDB) ReadStream(streamId storage.EventStreamId) ([]*storage.Event, error) {
	events := make([]*storage.Event, 0)
	session, err := mgo.Dial(store.baseUrl)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	collection := session.DB(store.database).C(store.collection)

	Log.Debug("Stream-ID: %v", streamId)
	err = collection.Find(bson.M{"eventid": streamId}).All(&events)
	Log.Debug("Events: %v", events)
	if err != nil {
		return nil, err
	}
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

	return events, nil
}
