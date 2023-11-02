// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"juice/public/ent/userfollowinfo"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserFollowInfo is the model entity for the UserFollowInfo schema.
type UserFollowInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint64 `json:"user_id,omitempty"`
	// FollowID holds the value of the "follow_id" field.
	FollowID uint64 `json:"follow_id,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   time.Time `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserFollowInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userfollowinfo.FieldID, userfollowinfo.FieldUserID, userfollowinfo.FieldFollowID, userfollowinfo.FieldStatus:
			values[i] = new(sql.NullInt64)
		case userfollowinfo.FieldCreateTime, userfollowinfo.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserFollowInfo fields.
func (ufi *UserFollowInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userfollowinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ufi.ID = int(value.Int64)
		case userfollowinfo.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ufi.UserID = uint64(value.Int64)
			}
		case userfollowinfo.FieldFollowID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field follow_id", values[i])
			} else if value.Valid {
				ufi.FollowID = uint64(value.Int64)
			}
		case userfollowinfo.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ufi.Status = int8(value.Int64)
			}
		case userfollowinfo.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ufi.CreateTime = value.Time
			}
		case userfollowinfo.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ufi.UpdateTime = value.Time
			}
		default:
			ufi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserFollowInfo.
// This includes values selected through modifiers, order, etc.
func (ufi *UserFollowInfo) Value(name string) (ent.Value, error) {
	return ufi.selectValues.Get(name)
}

// Update returns a builder for updating this UserFollowInfo.
// Note that you need to call UserFollowInfo.Unwrap() before calling this method if this UserFollowInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ufi *UserFollowInfo) Update() *UserFollowInfoUpdateOne {
	return NewUserFollowInfoClient(ufi.config).UpdateOne(ufi)
}

// Unwrap unwraps the UserFollowInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ufi *UserFollowInfo) Unwrap() *UserFollowInfo {
	_tx, ok := ufi.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserFollowInfo is not a transactional entity")
	}
	ufi.config.driver = _tx.drv
	return ufi
}

// String implements the fmt.Stringer.
func (ufi *UserFollowInfo) String() string {
	var builder strings.Builder
	builder.WriteString("UserFollowInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ufi.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ufi.UserID))
	builder.WriteString(", ")
	builder.WriteString("follow_id=")
	builder.WriteString(fmt.Sprintf("%v", ufi.FollowID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ufi.Status))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(ufi.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(ufi.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserFollowInfos is a parsable slice of UserFollowInfo.
type UserFollowInfos []*UserFollowInfo
