// Code generated by entc, DO NOT EDIT.

package book

import (
	"time"
)

const (
	// Label holds the string label denoting the book type in the database.
	Label = "book"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldPublishedDate holds the string denoting the published_date field in the database.
	FieldPublishedDate = "published_date"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeAuthors holds the string denoting the authors edge name in mutations.
	EdgeAuthors = "authors"
	// Table holds the table name of the book in the database.
	Table = "books"
	// AuthorsTable is the table that holds the authors relation/edge. The primary key declared below.
	AuthorsTable = "book_authors"
	// AuthorsInverseTable is the table name for the Author entity.
	// It exists in this package in order to avoid circular dependency with the "author" package.
	AuthorsInverseTable = "authors"
)

// Columns holds all SQL columns for book fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldPublishedDate,
	FieldImageURL,
	FieldDescription,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

var (
	// AuthorsPrimaryKey and AuthorsColumn2 are the table columns denoting the
	// primary key for the authors relation (M2M).
	AuthorsPrimaryKey = []string{"book_id", "author_id"}
)

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
