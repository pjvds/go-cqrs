package eventstore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/feeds"
	"github.com/pjvds/go-cqrs/sourcing"
	"net/http"
	"net/url"
)

type EventStore struct {
	baseUrl  string
	register *sourcing.EventTypeRegister
	PageSize int
}

func DailEventStore(url string, register *sourcing.EventTypeRegister) (*EventStore, error) {
	return &EventStore{
		baseUrl:  url,
		register: register,
		PageSize: 20,
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
			Data:      e,
		}
	}

	body, _ := json.Marshal(&data)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		Log.Error("Error while posting new stream request: %v", err)
		return err
	}

	if response.StatusCode != http.StatusCreated {
		Log.Error(fmt.Sprintf("Unexpected http status code in response: %v", response.Status))
		return errors.New(fmt.Sprintf("Unexpected http status code in response: %v", response.Status))
	}
	if location := response.Header.Get("location"); location != "" {
		Log.Notice(location)
	}

	return nil
}

type EventDocument struct {
}

func (store *EventStore) OpenStream(eventSourceId sourcing.EventSourceId) ([]*sourcing.EventEnvelope, error) {
	streamId := url.QueryEscape(eventSourceId.String())

	// Example: http://localhost:2113/streams/1b826790-5d4e-4227-7dc4-017ed73d30ac/head/backward/20
	url := fmt.Sprintf("%v/streams/%v/head/backward/%v", store.baseUrl, streamId, store.PageSize)

	feed, err := feeds.DownloadAtomFeed(url)
	if err != nil {
		return nil, err
	}

	result := make([]*sourcing.EventEnvelope, len(feed.Entries))
	for index, entry := range feed.Entries {
		alternateLink := entry.Links[1]
		eventUrl := alternateLink.Href

		eventEnvelope, err := downloadEvent(eventUrl)
		if err != nil {
			return nil, err
		}
		result[index] = eventEnvelope
	}

	return result, nil
}

func downloadEvent(url string) (*sourcing.EventEnvelope, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(response.Body)
	defer response.Body.Close()

	result := new(sourcing.EventEnvelope)
	err = decoder.Decode(result)

	return result, err
}
