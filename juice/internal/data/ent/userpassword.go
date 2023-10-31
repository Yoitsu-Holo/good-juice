// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"juice/internal/data/ent/userpassword"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserPassword is the model entity for the UserPassword schema.
type UserPassword struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// Salt holds the value of the "salt" field.
	Salt string `json:"salt,omitempty"`
	// Pwd holds the value of the "pwd" field.
	Pwd          string `json:"pwd,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserPassword) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userpassword.FieldID, userpassword.FieldUserID:
			values[i] = new(sql.NullInt64)
		case userpassword.FieldSalt, userpassword.FieldPwd:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserPassword fields.
func (up *UserPassword) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userpassword.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			up.ID = int(value.Int64)
		case userpassword.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				up.UserID = value.Int64
			}
		case userpassword.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				up.Salt = value.String
			}
		case userpassword.FieldPwd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pwd", values[i])
			} else if value.Valid {
				up.Pwd = value.String
			}
		default:
			up.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserPassword.
// This includes values selected through modifiers, order, etc.
func (up *UserPassword) Value(name string) (ent.Value, error) {
	return up.selectValues.Get(name)
}

// Update returns a builder for updating this UserPassword.
// Note that you need to call UserPassword.Unwrap() before calling this method if this UserPassword
// was returned from a transaction, and the transaction was committed or rolled back.
func (up *UserPassword) Update() *UserPasswordUpdateOne {
	return NewUserPasswordClient(up.config).UpdateOne(up)
}

// Unwrap unwraps the UserPassword entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (up *UserPassword) Unwrap() *UserPassword {
	_tx, ok := up.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserPassword is not a transactional entity")
	}
	up.config.driver = _tx.drv
	return up
}

// String implements the fmt.Stringer.
func (up *UserPassword) String() string {
	var builder strings.Builder
	builder.WriteString("UserPassword(")
	builder.WriteString(fmt.Sprintf("id=%v, ", up.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", up.UserID))
	builder.WriteString(", ")
	builder.WriteString("salt=")
	builder.WriteString(up.Salt)
	builder.WriteString(", ")
	builder.WriteString("pwd=")
	builder.WriteString(up.Pwd)
	builder.WriteByte(')')
	return builder.String()
}

// UserPasswords is a parsable slice of UserPassword.
type UserPasswords []*UserPassword
