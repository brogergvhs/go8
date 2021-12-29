// Code generated by entc, DO NOT EDIT.

package gen

import (
	"time"

	"github.com/gmhafiz/go8/ent/gen/author"
	"github.com/gmhafiz/go8/ent/gen/book"
	"github.com/gmhafiz/go8/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authorFields := schema.Author{}.Fields()
	_ = authorFields
	// authorDescCreatedAt is the schema descriptor for created_at field.
	authorDescCreatedAt := authorFields[4].Descriptor()
	// author.DefaultCreatedAt holds the default value on creation for the created_at field.
	author.DefaultCreatedAt = authorDescCreatedAt.Default.(func() time.Time)
	// authorDescUpdatedAt is the schema descriptor for updated_at field.
	authorDescUpdatedAt := authorFields[5].Descriptor()
	// author.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	author.DefaultUpdatedAt = authorDescUpdatedAt.Default.(func() time.Time)
	// author.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	author.UpdateDefaultUpdatedAt = authorDescUpdatedAt.UpdateDefault.(func() time.Time)
	bookFields := schema.Book{}.Fields()
	_ = bookFields
	// bookDescCreatedAt is the schema descriptor for created_at field.
	bookDescCreatedAt := bookFields[4].Descriptor()
	// book.DefaultCreatedAt holds the default value on creation for the created_at field.
	book.DefaultCreatedAt = bookDescCreatedAt.Default.(func() time.Time)
	// bookDescUpdatedAt is the schema descriptor for updated_at field.
	bookDescUpdatedAt := bookFields[5].Descriptor()
	// book.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	book.DefaultUpdatedAt = bookDescUpdatedAt.Default.(func() time.Time)
	// book.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	book.UpdateDefaultUpdatedAt = bookDescUpdatedAt.UpdateDefault.(func() time.Time)
}
