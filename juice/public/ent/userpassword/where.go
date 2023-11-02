// Code generated by ent, DO NOT EDIT.

package userpassword

import (
	"juice/public/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint64) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldUserID, v))
}

// Salt applies equality check predicate on the "salt" field. It's identical to SaltEQ.
func Salt(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldSalt, v))
}

// Pwd applies equality check predicate on the "pwd" field. It's identical to PwdEQ.
func Pwd(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldPwd, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldUpdateTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint64) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint64) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint64) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint64) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uint64) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uint64) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uint64) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uint64) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLTE(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.UserPassword {
	return predicate.UserPassword(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNotNull(FieldUserID))
}

// SaltEQ applies the EQ predicate on the "salt" field.
func SaltEQ(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldSalt, v))
}

// SaltNEQ applies the NEQ predicate on the "salt" field.
func SaltNEQ(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNEQ(FieldSalt, v))
}

// SaltIn applies the In predicate on the "salt" field.
func SaltIn(vs ...string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldIn(FieldSalt, vs...))
}

// SaltNotIn applies the NotIn predicate on the "salt" field.
func SaltNotIn(vs ...string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNotIn(FieldSalt, vs...))
}

// SaltGT applies the GT predicate on the "salt" field.
func SaltGT(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGT(FieldSalt, v))
}

// SaltGTE applies the GTE predicate on the "salt" field.
func SaltGTE(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGTE(FieldSalt, v))
}

// SaltLT applies the LT predicate on the "salt" field.
func SaltLT(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLT(FieldSalt, v))
}

// SaltLTE applies the LTE predicate on the "salt" field.
func SaltLTE(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLTE(FieldSalt, v))
}

// SaltContains applies the Contains predicate on the "salt" field.
func SaltContains(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldContains(FieldSalt, v))
}

// SaltHasPrefix applies the HasPrefix predicate on the "salt" field.
func SaltHasPrefix(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldHasPrefix(FieldSalt, v))
}

// SaltHasSuffix applies the HasSuffix predicate on the "salt" field.
func SaltHasSuffix(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldHasSuffix(FieldSalt, v))
}

// SaltEqualFold applies the EqualFold predicate on the "salt" field.
func SaltEqualFold(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEqualFold(FieldSalt, v))
}

// SaltContainsFold applies the ContainsFold predicate on the "salt" field.
func SaltContainsFold(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldContainsFold(FieldSalt, v))
}

// PwdEQ applies the EQ predicate on the "pwd" field.
func PwdEQ(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldPwd, v))
}

// PwdNEQ applies the NEQ predicate on the "pwd" field.
func PwdNEQ(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNEQ(FieldPwd, v))
}

// PwdIn applies the In predicate on the "pwd" field.
func PwdIn(vs ...string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldIn(FieldPwd, vs...))
}

// PwdNotIn applies the NotIn predicate on the "pwd" field.
func PwdNotIn(vs ...string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNotIn(FieldPwd, vs...))
}

// PwdGT applies the GT predicate on the "pwd" field.
func PwdGT(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGT(FieldPwd, v))
}

// PwdGTE applies the GTE predicate on the "pwd" field.
func PwdGTE(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGTE(FieldPwd, v))
}

// PwdLT applies the LT predicate on the "pwd" field.
func PwdLT(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLT(FieldPwd, v))
}

// PwdLTE applies the LTE predicate on the "pwd" field.
func PwdLTE(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLTE(FieldPwd, v))
}

// PwdContains applies the Contains predicate on the "pwd" field.
func PwdContains(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldContains(FieldPwd, v))
}

// PwdHasPrefix applies the HasPrefix predicate on the "pwd" field.
func PwdHasPrefix(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldHasPrefix(FieldPwd, v))
}

// PwdHasSuffix applies the HasSuffix predicate on the "pwd" field.
func PwdHasSuffix(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldHasSuffix(FieldPwd, v))
}

// PwdEqualFold applies the EqualFold predicate on the "pwd" field.
func PwdEqualFold(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEqualFold(FieldPwd, v))
}

// PwdContainsFold applies the ContainsFold predicate on the "pwd" field.
func PwdContainsFold(v string) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldContainsFold(FieldPwd, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.UserPassword {
	return predicate.UserPassword(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.UserPassword {
	return predicate.UserPassword(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.UserPassword {
	return predicate.UserPassword(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.UserPassword {
	return predicate.UserPassword(sql.FieldNotNull(FieldUpdateTime))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserPassword) predicate.UserPassword {
	return predicate.UserPassword(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserPassword) predicate.UserPassword {
	return predicate.UserPassword(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserPassword) predicate.UserPassword {
	return predicate.UserPassword(sql.NotPredicates(p))
}
