package retry

import (
	"context"
	"time"
)

// Do runs fn until success, ctx cancel, or attempts exhausted.
func Do(ctx context.Context, attempts int, backoff time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		err = fn()
		if err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(backoff):
		}
	}
	return err
}