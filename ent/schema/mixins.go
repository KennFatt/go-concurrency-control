package schema

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// VersionMixin implements `version` column with unsigned int type.
type VersionMixin struct {
	mixin.Schema
}

func (VersionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("version").Default(0).Optional(),
	}
}

func (VersionMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				switch m.Op() {
				case ent.OpCreate, ent.OpUpdate, ent.OpUpdateOne:
					nextVersion := time.Now().UnixNano()

					err := m.SetField("version", uint64(nextVersion))
					if err != nil {
						return nil, fmt.Errorf("versionmixin fail to set field version: %w", err)
					}
				}

				return next.Mutate(ctx, m)
			})
		},
	}
}
