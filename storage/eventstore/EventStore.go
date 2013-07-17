package eventstore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pjvds/go-cqrs/sourcing"
	"io/ioutil"
	"net/http"
	"net/url"
)

type EventStore struct {
	baseUrl string
}

func DailEventStore(url string) (*EventStore, error) {
	return &EventStore{
		baseUrl: url,
	}, nil
}

type Event struct {
	EventId   string      `json:"eventId"`
	EventType string      `json:"eventType"`
	Data      interface{} `json:"data"`
}

func (store *EventStore) NewStream(source sourcing.EventSource) error {
	streamId := url.QueryEscape(source.Id().String())
	url := fmt.Sprintf("%v/streams/%v", store.baseUrl, streamId)
	Log.Debug("Creating new stream at %v", url)

	events := source.Events()
	data := make([]Event, len(events))

	for i := 0; i < len(events); i++ {
		e := events[i]
		data[i] = Event{
			EventId:   e.EventId.String(),
			EventType: e.Name.String(),
			Data:      e.Payload,
		}
	}

	body, _ := json.Marshal(&data)
	Log.Debug("Posting: %v", string(body))
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		Log.Error("Error while posting new stream request: %v", err)
		return err
	}

	if response.StatusCode != http.StatusCreated {
		msg, _ := ioutil.ReadAll(response.Body)
		Log.Debug(string(msg))
		Log.Error(fmt.Sprintf("Unexpected http status code in response: %v", response.Status))
		return errors.New(fmt.Sprintf("Unexpected http status code in response: %v", response.Status))
	}

	return nil
}
