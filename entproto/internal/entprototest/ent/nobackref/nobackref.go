// Code generated by ent, DO NOT EDIT.

package nobackref

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the nobackref type in the database.
	Label = "no_backref"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// Table holds the table name of the nobackref in the database.
	Table = "no_backrefs"
	// ImagesTable is the table that holds the images relation/edge.
	ImagesTable = "images"
	// ImagesInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ImagesInverseTable = "images"
	// ImagesColumn is the table column denoting the images relation/edge.
	ImagesColumn = "no_backref_images"
)

// Columns holds all SQL columns for nobackref fields.
var Columns = []string{
	FieldID,
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

// Order defines the ordering method for the NoBackref queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByImagesCount orders the results by images count.
func ByImagesCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newImagesStep(), opts...)
	}
}

// ByImages orders the results by images terms.
func ByImages(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newImagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
	)
}
