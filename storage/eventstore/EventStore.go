package eventstore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pjvds/feeds"
	"github.com/pjvds/go-cqrs/storage"
	"github.com/pjvds/go-cqrs/storage/serialization"
	"io/ioutil"
	"net/http"
	"net/url"
)

type EventStore struct {
	baseUrl    string
	PageSize   int
	serializer serialization.Serializer
}

func DailEventStore(url string, register *serialization.EventTypeRegister) (*EventStore, error) {
	return &EventStore{
		baseUrl:    url,
		PageSize:   20,
		serializer: serialization.NewJsonSerializer(register),
	}, nil
}

type Event struct {
	EventId   string         `json:"eventId"`
	EventType string         `json:"eventType"`
	Data      *storage.Event `json:"data"`
}

func (store *EventStore) WriteStream(change *storage.EventStreamChange) error {
	streamId := url.QueryEscape(change.StreamId.String())
	url := fmt.Sprintf("%v/streams/%v", store.baseUrl, streamId)

	events := change.Events
	data := make([]*Event, len(events))
	Log.Debug("Creating new stream at %v for %v events", url, len(events))

	for i := 0; i < len(events); i++ {
		e := events[i]

		data[i] = &Event{
			EventId:   e.EventId.String(),
			EventType: e.Name.String(),
			Data:      e,
		}
	}

	body, _ := json.MarshalIndent(&data, "", "  ")
	Log.Debug("Posting:\n%v", string(body))

	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		Log.Error("Error while posting new stream request: %v\n\tbody: %v", err, string(body))
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

func (store *EventStore) ReadStream(streamId storage.EventStreamId) ([]*storage.Event, error) {
	events := make([]*storage.Event, 0)
	pointer, err := OpenStreamPointer(streamId.String(), store.PageSize)
	if err != nil {
		return nil, err
	}

	for pointer != nil {
		event, err := store.downloadEvent(pointer.EventUrl)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
		next, err := pointer.Next()

		pointer = next
	}

	return events, nil
}

func linksToMap(links []*feeds.AtomLink) map[string]string {
	m := make(map[string]string, len(links))
	for _, link := range links {
		m[link.Rel] = link.Href
	}

	return m
}

func (store *EventStore) processFeed(feed *feeds.AtomFeed) ([]*storage.Event, error) {
	result := make([]*storage.Event, len(feed.Entries))
	for index, entry := range feed.Entries {
		alternateLink := entry.Links[1]
		eventUrl := alternateLink.Href

		event, err := store.downloadEvent(eventUrl)
		if err != nil {
			return nil, err
		}
		result[index] = event
	}

	return result, nil
}

func (store *EventStore) downloadEvent(url string) (*storage.Event, error) {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Accept", "application/json")
	c := http.Client{}

	response, err := c.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	event, err := store.serializer.Deserialize(*storage.NewEventName(""), body)
	return event, err
}
