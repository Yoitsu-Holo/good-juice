// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"juice/app/user/internal/data/ent/videolike"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// VideoLike is the model entity for the VideoLike schema.
type VideoLike struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// VideoID holds the value of the "video_id" field.
	VideoID int `json:"video_id,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   time.Time `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VideoLike) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case videolike.FieldID, videolike.FieldUserID, videolike.FieldVideoID, videolike.FieldStatus:
			values[i] = new(sql.NullInt64)
		case videolike.FieldCreateTime, videolike.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VideoLike fields.
func (vl *VideoLike) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case videolike.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vl.ID = int(value.Int64)
		case videolike.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				vl.UserID = int(value.Int64)
			}
		case videolike.FieldVideoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field video_id", values[i])
			} else if value.Valid {
				vl.VideoID = int(value.Int64)
			}
		case videolike.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				vl.Status = int8(value.Int64)
			}
		case videolike.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				vl.CreateTime = value.Time
			}
		case videolike.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				vl.UpdateTime = value.Time
			}
		default:
			vl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VideoLike.
// This includes values selected through modifiers, order, etc.
func (vl *VideoLike) Value(name string) (ent.Value, error) {
	return vl.selectValues.Get(name)
}

// Update returns a builder for updating this VideoLike.
// Note that you need to call VideoLike.Unwrap() before calling this method if this VideoLike
// was returned from a transaction, and the transaction was committed or rolled back.
func (vl *VideoLike) Update() *VideoLikeUpdateOne {
	return NewVideoLikeClient(vl.config).UpdateOne(vl)
}

// Unwrap unwraps the VideoLike entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vl *VideoLike) Unwrap() *VideoLike {
	_tx, ok := vl.config.driver.(*txDriver)
	if !ok {
		panic("ent: VideoLike is not a transactional entity")
	}
	vl.config.driver = _tx.drv
	return vl
}

// String implements the fmt.Stringer.
func (vl *VideoLike) String() string {
	var builder strings.Builder
	builder.WriteString("VideoLike(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vl.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", vl.UserID))
	builder.WriteString(", ")
	builder.WriteString("video_id=")
	builder.WriteString(fmt.Sprintf("%v", vl.VideoID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", vl.Status))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(vl.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(vl.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// VideoLikes is a parsable slice of VideoLike.
type VideoLikes []*VideoLike