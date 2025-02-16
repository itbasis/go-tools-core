package option

import (
	"context"
	"log/slog"
	"slices"
)

type RestoreOption[T any] interface {
	Option[T]

	Save(ctx context.Context, obj *T) error
	Restore(ctx context.Context, obj *T) error
}

func ApplyRestoreOptions[T any](ctx context.Context, obj *T, opts []RestoreOption[T], action func()) error {
	var keys = make(map[Key]struct{}, len(opts))

	for _, opt := range opts {
		var key = opt.Key()

		if err := _existKey(keys, key); err != nil {
			slog.Error(_msgAlreadyBefore, _slogAttrOptionKey(key))

			return err
		}

		keys[key] = struct{}{}

		if err := opt.Save(ctx, obj); err != nil {
			return err //nolint:wrapcheck // TODO
		}

		if err := opt.Apply(ctx, obj); err != nil {
			return err //nolint:wrapcheck // TODO
		}
	}

	action()

	for _, opt := range slices.Backward(opts) {
		if err := opt.Restore(ctx, obj); err != nil {
			return err //nolint:wrapcheck // TODO
		}
	}

	return nil
}
