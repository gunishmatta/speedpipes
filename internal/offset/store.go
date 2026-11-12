package offset

import "context"

type Store interface {
	Load(ctx context.Context) (int64, error)
	Save(ctx context.Context, offset int64) error
}
