// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-cc/ent/seat"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Seat is the model entity for the Seat schema.
type Seat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IsBooked holds the value of the "is_booked" field.
	IsBooked bool `json:"is_booked,omitempty"`
	// PassangerName holds the value of the "passanger_name" field.
	PassangerName *string `json:"passanger_name,omitempty"`
	// Version holds the value of the "version" field.
	Version      uint64 `json:"version,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Seat) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case seat.FieldIsBooked:
			values[i] = new(sql.NullBool)
		case seat.FieldID, seat.FieldVersion:
			values[i] = new(sql.NullInt64)
		case seat.FieldPassangerName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Seat fields.
func (s *Seat) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case seat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case seat.FieldIsBooked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_booked", values[i])
			} else if value.Valid {
				s.IsBooked = value.Bool
			}
		case seat.FieldPassangerName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field passanger_name", values[i])
			} else if value.Valid {
				s.PassangerName = new(string)
				*s.PassangerName = value.String
			}
		case seat.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				s.Version = uint64(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Seat.
// This includes values selected through modifiers, order, etc.
func (s *Seat) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Seat.
// Note that you need to call Seat.Unwrap() before calling this method if this Seat
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Seat) Update() *SeatUpdateOne {
	return NewSeatClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Seat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Seat) Unwrap() *Seat {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Seat is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Seat) String() string {
	var builder strings.Builder
	builder.WriteString("Seat(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("is_booked=")
	builder.WriteString(fmt.Sprintf("%v", s.IsBooked))
	builder.WriteString(", ")
	if v := s.PassangerName; v != nil {
		builder.WriteString("passanger_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", s.Version))
	builder.WriteByte(')')
	return builder.String()
}

// Seats is a parsable slice of Seat.
type Seats []*Seat