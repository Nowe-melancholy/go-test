// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-test/ent/middleware"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Middleware is the model entity for the Middleware schema.
type Middleware struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// LID holds the value of the "l_id" field.
	LID int `json:"l_id,omitempty"`
	// DID holds the value of the "d_id" field.
	DID int `json:"d_id,omitempty"`
	// SysID holds the value of the "sys_id" field.
	SysID int `json:"sys_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Middleware) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case middleware.FieldID, middleware.FieldLID, middleware.FieldDID, middleware.FieldSysID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Middleware", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Middleware fields.
func (m *Middleware) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case middleware.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case middleware.FieldLID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field l_id", values[i])
			} else if value.Valid {
				m.LID = int(value.Int64)
			}
		case middleware.FieldDID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field d_id", values[i])
			} else if value.Valid {
				m.DID = int(value.Int64)
			}
		case middleware.FieldSysID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sys_id", values[i])
			} else if value.Valid {
				m.SysID = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Middleware.
// Note that you need to call Middleware.Unwrap() before calling this method if this Middleware
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Middleware) Update() *MiddlewareUpdateOne {
	return (&MiddlewareClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Middleware entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Middleware) Unwrap() *Middleware {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Middleware is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Middleware) String() string {
	var builder strings.Builder
	builder.WriteString("Middleware(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("l_id=")
	builder.WriteString(fmt.Sprintf("%v", m.LID))
	builder.WriteString(", ")
	builder.WriteString("d_id=")
	builder.WriteString(fmt.Sprintf("%v", m.DID))
	builder.WriteString(", ")
	builder.WriteString("sys_id=")
	builder.WriteString(fmt.Sprintf("%v", m.SysID))
	builder.WriteByte(')')
	return builder.String()
}

// Middlewares is a parsable slice of Middleware.
type Middlewares []*Middleware

func (m Middlewares) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
