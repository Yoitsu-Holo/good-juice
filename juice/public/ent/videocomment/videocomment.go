// Code generated by ent, DO NOT EDIT.

package videocomment

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the videocomment type in the database.
	Label = "video_comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCommentID holds the string denoting the comment_id field in the database.
	FieldCommentID = "comment_id"
	// FieldPcommentID holds the string denoting the pcomment_id field in the database.
	FieldPcommentID = "pcomment_id"
	// FieldVideoID holds the string denoting the video_id field in the database.
	FieldVideoID = "video_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCommentText holds the string denoting the comment_text field in the database.
	FieldCommentText = "comment_text"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the videocomment in the database.
	Table = "video_comment"
)

// Columns holds all SQL columns for videocomment fields.
var Columns = []string{
	FieldID,
	FieldCommentID,
	FieldPcommentID,
	FieldVideoID,
	FieldUserID,
	FieldCommentText,
	FieldCreateTime,
	FieldUpdateTime,
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

// OrderOption defines the ordering options for the VideoComment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCommentID orders the results by the comment_id field.
func ByCommentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommentID, opts...).ToFunc()
}

// ByPcommentID orders the results by the pcomment_id field.
func ByPcommentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPcommentID, opts...).ToFunc()
}

// ByVideoID orders the results by the video_id field.
func ByVideoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVideoID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByCommentText orders the results by the comment_text field.
func ByCommentText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommentText, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}
