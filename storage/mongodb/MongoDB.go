package mongodb

import (
	//"bytes"
	//"encoding/json"
	. "github.com/dominikmayer/go-cqrs/storage"
	"github.com/dominikmayer/go-cqrs/storage/serialization"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type MongoDB struct {
	baseUrl    string
	database   string
	collection string
}

func DialMongoDB(url string, database string, collection string, register *serialization.EventTypeRegister) (*MongoDB, error) {
	return &MongoDB{
		baseUrl:    url,
		database:   database,
		collection: collection,
	}, nil
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
	//marshallid = NewEventStreamIdFromString(streamId)
	//marshallid = []byte(streamId.String())
	Log.Debug("Stream-ID: %v, Error: %v", marshallid, marshallerr)
	mybson := bson.M{"streamid": marshallid}//bson.Raw{Kind: 0, Data: marshallid,}}
	Log.Debug("BSON: %v", mybson)
	err = collection.Find(mybson).All(&events)
	//err = collection.Find(bson.M{"events.name": "github.com/dominikmayer/go-cqrs/tests/events/UsernameChanged"}).All(&events)
	Log.Debug("%v Events: %v", len(events), events)
	if err != nil {
		return nil, err
	}
	//events = events.Events
	receivedEvents := events[0].Events
	Log.Debug("%v Events: %v", len(receivedEvents), receivedEvents)
	//err = collection.FindId(streamId).All(&events)


	return receivedEvents, nil
}
