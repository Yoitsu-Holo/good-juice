// Code generated by ent, DO NOT EDIT.

package videolike

import (
	"juice/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldUserID, v))
}

// VideoID applies equality check predicate on the "video_id" field. It's identical to VideoIDEQ.
func VideoID(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldVideoID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldStatus, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldUpdateTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLTE(FieldUserID, v))
}

// VideoIDEQ applies the EQ predicate on the "video_id" field.
func VideoIDEQ(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldVideoID, v))
}

// VideoIDNEQ applies the NEQ predicate on the "video_id" field.
func VideoIDNEQ(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNEQ(FieldVideoID, v))
}

// VideoIDIn applies the In predicate on the "video_id" field.
func VideoIDIn(vs ...int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldIn(FieldVideoID, vs...))
}

// VideoIDNotIn applies the NotIn predicate on the "video_id" field.
func VideoIDNotIn(vs ...int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNotIn(FieldVideoID, vs...))
}

// VideoIDGT applies the GT predicate on the "video_id" field.
func VideoIDGT(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGT(FieldVideoID, v))
}

// VideoIDGTE applies the GTE predicate on the "video_id" field.
func VideoIDGTE(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGTE(FieldVideoID, v))
}

// VideoIDLT applies the LT predicate on the "video_id" field.
func VideoIDLT(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLT(FieldVideoID, v))
}

// VideoIDLTE applies the LTE predicate on the "video_id" field.
func VideoIDLTE(v int64) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLTE(FieldVideoID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLTE(FieldStatus, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.VideoLike {
	return predicate.VideoLike(sql.FieldLTE(FieldUpdateTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VideoLike) predicate.VideoLike {
	return predicate.VideoLike(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VideoLike) predicate.VideoLike {
	return predicate.VideoLike(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VideoLike) predicate.VideoLike {
	return predicate.VideoLike(sql.NotPredicates(p))
}
