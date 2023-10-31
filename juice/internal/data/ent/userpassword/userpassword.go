// Code generated by ent, DO NOT EDIT.

package userpassword

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the userpassword type in the database.
	Label = "user_password"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldPwd holds the string denoting the pwd field in the database.
	FieldPwd = "pwd"
	// Table holds the table name of the userpassword in the database.
	Table = "user_passwords"
)

// Columns holds all SQL columns for userpassword fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldSalt,
	FieldPwd,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(int64) error
)

// OrderOption defines the ordering options for the UserPassword queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// BySalt orders the results by the salt field.
func BySalt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSalt, opts...).ToFunc()
}

// ByPwd orders the results by the pwd field.
func ByPwd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPwd, opts...).ToFunc()
}
