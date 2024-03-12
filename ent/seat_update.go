// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-cc/ent/predicate"
	"go-cc/ent/seat"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SeatUpdate is the builder for updating Seat entities.
type SeatUpdate struct {
	config
	hooks    []Hook
	mutation *SeatMutation
}

// Where appends a list predicates to the SeatUpdate builder.
func (su *SeatUpdate) Where(ps ...predicate.Seat) *SeatUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetIsBooked sets the "is_booked" field.
func (su *SeatUpdate) SetIsBooked(b bool) *SeatUpdate {
	su.mutation.SetIsBooked(b)
	return su
}

// SetNillableIsBooked sets the "is_booked" field if the given value is not nil.
func (su *SeatUpdate) SetNillableIsBooked(b *bool) *SeatUpdate {
	if b != nil {
		su.SetIsBooked(*b)
	}
	return su
}

// SetPassengerName sets the "passenger_name" field.
func (su *SeatUpdate) SetPassengerName(s string) *SeatUpdate {
	su.mutation.SetPassengerName(s)
	return su
}

// SetNillablePassengerName sets the "passenger_name" field if the given value is not nil.
func (su *SeatUpdate) SetNillablePassengerName(s *string) *SeatUpdate {
	if s != nil {
		su.SetPassengerName(*s)
	}
	return su
}

// ClearPassengerName clears the value of the "passenger_name" field.
func (su *SeatUpdate) ClearPassengerName() *SeatUpdate {
	su.mutation.ClearPassengerName()
	return su
}

// SetVersion sets the "version" field.
func (su *SeatUpdate) SetVersion(u uint64) *SeatUpdate {
	su.mutation.ResetVersion()
	su.mutation.SetVersion(u)
	return su
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (su *SeatUpdate) SetNillableVersion(u *uint64) *SeatUpdate {
	if u != nil {
		su.SetVersion(*u)
	}
	return su
}

// AddVersion adds u to the "version" field.
func (su *SeatUpdate) AddVersion(u int64) *SeatUpdate {
	su.mutation.AddVersion(u)
	return su
}

// ClearVersion clears the value of the "version" field.
func (su *SeatUpdate) ClearVersion() *SeatUpdate {
	su.mutation.ClearVersion()
	return su
}

// Mutation returns the SeatMutation object of the builder.
func (su *SeatUpdate) Mutation() *SeatMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SeatUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SeatUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SeatUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SeatUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SeatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(seat.Table, seat.Columns, sqlgraph.NewFieldSpec(seat.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.IsBooked(); ok {
		_spec.SetField(seat.FieldIsBooked, field.TypeBool, value)
	}
	if value, ok := su.mutation.PassengerName(); ok {
		_spec.SetField(seat.FieldPassengerName, field.TypeString, value)
	}
	if su.mutation.PassengerNameCleared() {
		_spec.ClearField(seat.FieldPassengerName, field.TypeString)
	}
	if value, ok := su.mutation.Version(); ok {
		_spec.SetField(seat.FieldVersion, field.TypeUint64, value)
	}
	if value, ok := su.mutation.AddedVersion(); ok {
		_spec.AddField(seat.FieldVersion, field.TypeUint64, value)
	}
	if su.mutation.VersionCleared() {
		_spec.ClearField(seat.FieldVersion, field.TypeUint64)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{seat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SeatUpdateOne is the builder for updating a single Seat entity.
type SeatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SeatMutation
}

// SetIsBooked sets the "is_booked" field.
func (suo *SeatUpdateOne) SetIsBooked(b bool) *SeatUpdateOne {
	suo.mutation.SetIsBooked(b)
	return suo
}

// SetNillableIsBooked sets the "is_booked" field if the given value is not nil.
func (suo *SeatUpdateOne) SetNillableIsBooked(b *bool) *SeatUpdateOne {
	if b != nil {
		suo.SetIsBooked(*b)
	}
	return suo
}

// SetPassengerName sets the "passenger_name" field.
func (suo *SeatUpdateOne) SetPassengerName(s string) *SeatUpdateOne {
	suo.mutation.SetPassengerName(s)
	return suo
}

// SetNillablePassengerName sets the "passenger_name" field if the given value is not nil.
func (suo *SeatUpdateOne) SetNillablePassengerName(s *string) *SeatUpdateOne {
	if s != nil {
		suo.SetPassengerName(*s)
	}
	return suo
}

// ClearPassengerName clears the value of the "passenger_name" field.
func (suo *SeatUpdateOne) ClearPassengerName() *SeatUpdateOne {
	suo.mutation.ClearPassengerName()
	return suo
}

// SetVersion sets the "version" field.
func (suo *SeatUpdateOne) SetVersion(u uint64) *SeatUpdateOne {
	suo.mutation.ResetVersion()
	suo.mutation.SetVersion(u)
	return suo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (suo *SeatUpdateOne) SetNillableVersion(u *uint64) *SeatUpdateOne {
	if u != nil {
		suo.SetVersion(*u)
	}
	return suo
}

// AddVersion adds u to the "version" field.
func (suo *SeatUpdateOne) AddVersion(u int64) *SeatUpdateOne {
	suo.mutation.AddVersion(u)
	return suo
}

// ClearVersion clears the value of the "version" field.
func (suo *SeatUpdateOne) ClearVersion() *SeatUpdateOne {
	suo.mutation.ClearVersion()
	return suo
}

// Mutation returns the SeatMutation object of the builder.
func (suo *SeatUpdateOne) Mutation() *SeatMutation {
	return suo.mutation
}

// Where appends a list predicates to the SeatUpdate builder.
func (suo *SeatUpdateOne) Where(ps ...predicate.Seat) *SeatUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SeatUpdateOne) Select(field string, fields ...string) *SeatUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Seat entity.
func (suo *SeatUpdateOne) Save(ctx context.Context) (*Seat, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SeatUpdateOne) SaveX(ctx context.Context) *Seat {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SeatUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SeatUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SeatUpdateOne) sqlSave(ctx context.Context) (_node *Seat, err error) {
	_spec := sqlgraph.NewUpdateSpec(seat.Table, seat.Columns, sqlgraph.NewFieldSpec(seat.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Seat.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, seat.FieldID)
		for _, f := range fields {
			if !seat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != seat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.IsBooked(); ok {
		_spec.SetField(seat.FieldIsBooked, field.TypeBool, value)
	}
	if value, ok := suo.mutation.PassengerName(); ok {
		_spec.SetField(seat.FieldPassengerName, field.TypeString, value)
	}
	if suo.mutation.PassengerNameCleared() {
		_spec.ClearField(seat.FieldPassengerName, field.TypeString)
	}
	if value, ok := suo.mutation.Version(); ok {
		_spec.SetField(seat.FieldVersion, field.TypeUint64, value)
	}
	if value, ok := suo.mutation.AddedVersion(); ok {
		_spec.AddField(seat.FieldVersion, field.TypeUint64, value)
	}
	if suo.mutation.VersionCleared() {
		_spec.ClearField(seat.FieldVersion, field.TypeUint64)
	}
	_node = &Seat{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{seat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
