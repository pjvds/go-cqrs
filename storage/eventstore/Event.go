package eventstore

import (
	"fmt"
	"github.com/pjvds/feeds"
)

type StreamEventPointer struct {
	EventUrl string

	entryIndex int
	pageSize   int
	pageUrl    string

	page *feeds.AtomFeed
}

func OpenStreamPointer(streamId string, pageSize int) (*StreamEventPointer, error) {
	url := fmt.Sprintf("http://localhost:2113/streams/%v/0/forward/%v", streamId, pageSize)
	feed, err := feeds.DownloadAtomFeed(url)
	if err != nil {
		return nil, err
	}

	firstEntry := feed.Entries[0]
	link, _ := firstEntry.Link("alternate")
	return NewStreamEventPointer(url, feed, 0, pageSize, link), nil
}

func NewStreamEventPointer(pageUrl string, page *feeds.AtomFeed, entryIndex int, pageSize int, url string) *StreamEventPointer {
	return &StreamEventPointer{
		pageUrl:    pageUrl,
		page:       page,
		entryIndex: entryIndex,
		pageSize:   pageSize,
		EventUrl:   url,
	}
}

func (s *StreamEventPointer) Next() (*StreamEventPointer, error) {
	Log.Debug("Looking for next event for %v", s.EventUrl)

	// when this this not the last, get next event on page
	if !s.isLastEvent() {
		entryIndex := s.entryIndex + 1

		// check if we have a next event on the current page
		if entryIndex >= len(s.page.Entries) {
			// current event is the last one, download feed to see if
			// it has new events.
			Log.Debug("Event is last known event on page, looking for more events in page")
			page, err := feeds.DownloadAtomFeed(s.pageUrl)
			if err != nil {
				return nil, err
			}
			s.page = page

			// check if there are new events on the feed
			if entryIndex >= len(s.page.Entries) {
				Log.Debug("No new event in page, current event is the last")
				return nil, nil
			}
			Log.Debug("New events in page after downloaded")
		}

		entry := s.page.Entries[entryIndex]
		link, _ := entry.Link("alternate")
		pointer := NewStreamEventPointer(s.pageUrl, s.page, entryIndex, s.pageSize, link)
		return pointer, nil
	}

	// check if current feed has next link, otherwise redownload it
	if _, ok := s.page.Link("next"); !ok {
		page, err := feeds.DownloadAtomFeed(s.pageUrl)
		if err != nil {
			return nil, err
		}

		s.page = page
	}

	if link, ok := s.page.Link("next"); ok {
		nextPage, err := feeds.DownloadAtomFeed(link)
		if err != nil {
			return nil, err
		}

		entryIndex := 0
		entry := nextPage.Entries[entryIndex]
		link, _ := entry.Link("alternate")
		pointer := NewStreamEventPointer(link, nextPage, entryIndex, s.pageSize, link)
		return pointer, nil
	}
	return nil, nil
}

func (s *StreamEventPointer) isLastEvent() bool {
	return s.entryIndex+1 == s.pageSize
}
