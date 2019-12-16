package rethinkdb

import (
	. "github.com/pjvds/go-cqrs/storage"
	"github.com/pjvds/go-cqrs/storage/serialization"
	r "github.com/dancannon/gorethink"
)

type RethinkDB struct {
	address    string
	database   string
	table	   string
}

func New(url string, database string, table string, register *serialization.EventTypeRegister) *RethinkDB {
  	return &RethinkDB{
		address:    url,  // TODO: os.Getenv("RETHINKDB_URL")
		database:   database,
		table: table,
	}
}

func (store *RethinkDB) WriteStream(change *EventStreamChange) error {
	session, err := r.Connect(map[string]interface{}{
		"address": store.address,
		"database": store.database,
})
	if err != nil {
		return err
	}
	defer session.Close()

	_, err = r.Table(store.table).Insert(change.GetPersistableObject()).RunWrite(session)
	Log.Debug("Change: %v, Error: %v", change, err)
	return err
}

func (store *RethinkDB) ReadStream(streamId EventStreamId) ([]*Event, error) {
	session, err := r.Connect(map[string]interface{}{
	"address": store.address,
	"database": store.database,
})
	if err != nil {
		return nil, err
	}
	defer session.Close()

	row, err := r.Table(store.table).Filter(streamId).RunRow(session)

	Log.Debug("Stream-ID: %v", streamId)
	if err != nil {
		Log.Debug("Error: %v", err)
		return nil, err
	}

	var persistedEvent EventStreamChangePersist
	err = row.Scan(&persistedEvent)
	if err != nil {
		Log.Debug("Error: %v", err)
		return nil, err
	}

	receivedEvents := persistedEvent.Events
	Log.Debug("%v unpacked Events: %v", len(receivedEvents), receivedEvents)

	return receivedEvents, nil
}
