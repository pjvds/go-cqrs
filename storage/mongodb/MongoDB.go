package mongodb

import (
	. "github.com/dominikmayer/go-cqrs/storage"
	"github.com/dominikmayer/go-cqrs/storage/serialization"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"errors"
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

func (store *MongoDB) WriteStream(change *EventStreamChange) error {
	session, err := mgo.Dial(store.baseUrl)
	//Log.Debug("Base-URL: %v", store.baseUrl)
	if err != nil {
		return err
	}
	defer session.Close()

	collection := session.DB(store.database).C(store.collection)

	//Log.Debug("Inserting data: %v", change)
	err = collection.Insert(change.GetPersistableObject())
	if err != nil {
		return err
	}

	return nil
}

func (store *MongoDB) ReadStream(streamId EventStreamId) ([]*Event, error) {
	persistedEvents := make([]*EventStreamChangePersist,0)
	session, err := mgo.Dial(store.baseUrl)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	collection := session.DB(store.database).C(store.collection)

	Log.Debug("Stream-ID: %v", streamId)
	mybson := bson.M{"streamid": streamId.String()}
	err = collection.Find(mybson).All(&persistedEvents)
	if err != nil {
		Log.Debug("Error: %v", err)
		return nil, err
	}
	
	numberOfEvents := len(persistedEvents)
	if numberOfEvents > 1 {
		Log.Debug("%v duplicate objects: %v", numberOfEvents, persistedEvents)
		return nil, errors.New("Duplicate objects found")
	}

	receivedEvents := persistedEvents[0].Events
	Log.Debug("%v unpacked Events: %v", len(receivedEvents), receivedEvents)

	return receivedEvents, nil
}
