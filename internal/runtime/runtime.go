package runtime

import (
	"context"

	"github.com/gunishmatta/speedpipes/internal/offset"
	"github.com/gunishmatta/speedpipes/internal/source"
)

type Runtime struct {
	source source.Source
	store  offset.FileStore
}

func New(src source.Source, offsetStore offset.FileStore) *Runtime {
	return &Runtime{
		source: src,
		store:  offsetStore,
	}
}

func (r *Runtime) Run(ctx context.Context) error {
	currentOffset, err := r.store.Load(ctx)
	if err != nil {
		return err
	}
	events, err := r.source.Poll(ctx, currentOffset)
	if err != nil {
		return err
	}
	for _, event := range events {
		// Process the event (for now, we just print it)
		println("Processing event ID:", event.ID, "Value:", string(event.Value))
		if err := r.store.Save(ctx, event.ID); err != nil {
			return err
		}
	}
	return nil
}
