// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldInterests holds the string denoting the interests field in the database.
	FieldInterests = "interests"
	// FieldRating holds the string denoting the rating field in the database.
	FieldRating = "rating"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldAge,
	FieldDescription,
	FieldInterests,
	FieldRating,
	FieldImageURL,
	FieldTags,
	FieldPassword,
	FieldType,
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
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// InterestsValidator is a validator for the "interests" field. It is called by the builders before save.
	InterestsValidator func(string) error
	// RatingValidator is a validator for the "rating" field. It is called by the builders before save.
	RatingValidator func(int) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByInterests orders the results by the interests field.
func ByInterests(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInterests, opts...).ToFunc()
}

// ByRating orders the results by the rating field.
func ByRating(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRating, opts...).ToFunc()
}

// ByImageURL orders the results by the image_url field.
func ByImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImageURL, opts...).ToFunc()
}

// ByTags orders the results by the tags field.
func ByTags(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTags, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}
