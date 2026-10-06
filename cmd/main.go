package main

import (
	"context"
	"log"

	"github.com/gunishmatta/speedpipes/internal/runtime"
	"github.com/gunishmatta/speedpipes/internal/source"
)

func main() {
	events := []source.Event{
		{ID: 1, Value: []byte("Event 1")},
		{ID: 2, Value: []byte("Event 2")},
		{ID: 3, Value: []byte("Event 3")},
	}
	source := source.NewMemorySource(events)
	runtime := runtime.New(source)
	ctx := context.Background()
	if err := runtime.Run(ctx); err != nil {
		log.Fatalf("Error running runtime: %v", err)
	}
}
