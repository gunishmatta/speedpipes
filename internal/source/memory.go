package source

import "context"

type Event struct {
	ID    int64
	Value []byte
}

type Source interface {
	Poll(ctx context.Context, after int64) ([]Event, error)
}

type MemorySource struct {
	events []Event
}

func NewMemorySource(events []Event) *MemorySource {
	return &MemorySource{
		events: events,
	}
}

func (ms *MemorySource) Poll(ctx context.Context, after int64) ([]Event, error) {
	var result []Event
	for _, event := range ms.events {
		if event.ID > after {
			result = append(result, event)
		}
	}
	return result, nil
}
