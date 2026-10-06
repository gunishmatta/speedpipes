package runtime

import (
	"context"

	"github.com/gunishmatta/speedpipes/internal/source"
)

type Runtime struct {
	source source.Source
	offset int64
}

func New(src source.Source) *Runtime {
	return &Runtime{
		source: src,
		offset: 0,
	}
}

func (r *Runtime) Run(ctx context.Context) error {
	events, err := r.source.Poll(ctx, r.offset)
	if err != nil {
		return err
	}
	for _, event := range events {
		// Process the event (for now, we just print it)
		println("Processing event ID:", event.ID, "Value:", string(event.Value))
		r.offset = event.ID
	}
	return nil
}
