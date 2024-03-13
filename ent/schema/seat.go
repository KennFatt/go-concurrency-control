package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Seat holds the schema definition for the Seat entity.
type Seat struct {
	ent.Schema
}

// Fields of the Seat.
func (Seat) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Bool("is_booked").Default(false),
		field.String("passenger_name").Optional().Nillable(),
	}
}

// Edges of the Seat.
func (Seat) Edges() []ent.Edge {
	return nil
}

func (Seat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		VersionMixin{},
	}
}
