// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"test/ent/todo"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Item holds the value of the "item" field.
	Item string `json:"item,omitempty"`
	// Done holds the value of the "done" field.
	Done bool `json:"done,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case todo.FieldDone:
			values[i] = new(sql.NullBool)
		case todo.FieldID:
			values[i] = new(sql.NullInt64)
		case todo.FieldItem:
			values[i] = new(sql.NullString)
		case todo.FieldCreatedAt, todo.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case todo.FieldItem:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item", values[i])
			} else if value.Valid {
				t.Item = value.String
			}
		case todo.FieldDone:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field done", values[i])
			} else if value.Valid {
				t.Done = value.Bool
			}
		case todo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case todo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Todo.
// This includes values selected through modifiers, order, etc.
func (t *Todo) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Todo.
// Note that you need to call Todo.Unwrap() before calling this method if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return NewTodoClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Todo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("item=")
	builder.WriteString(t.Item)
	builder.WriteString(", ")
	builder.WriteString("done=")
	builder.WriteString(fmt.Sprintf("%v", t.Done))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Todos is a parsable slice of Todo.
type Todos []*Todo
