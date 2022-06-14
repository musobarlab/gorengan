package schema

import (
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/musobarlab/gorengan/internal/modules/category/domain"
)

// CategorySchemaInput struct
type CategorySchemaInput struct {
	ID           string
	Name         string
	CreatorID    *string
	CreatorIP    *string
	Created      *time.Time
	EditorID     *string
	EditorIP     *string
	LastModified *time.Time
	IsDeleted    *bool
	Deleted      *time.Time
}

// CategorySchema resolver
type CategorySchema struct {
	Category *domain.Category
}

// ID function
func (c *CategorySchema) ID() graphql.ID {
	return graphql.ID(c.Category.ID)
}

// Name function
func (c *CategorySchema) Name() string {
	return c.Category.Name
}

// CreatorID function
func (c *CategorySchema) CreatorID() string {
	return c.Category.CreatorID
}

// CreatorIP function
func (c *CategorySchema) CreatorIP() string {
	return c.Category.CreatorIP
}

// Created function
func (c *CategorySchema) Created() *graphql.Time {
	return &graphql.Time{Time: c.Category.Created}
}

// EditorID function
func (c *CategorySchema) EditorID() string {
	return c.Category.EditorID
}

// EditorIP function
func (c *CategorySchema) EditorIP() string {
	return c.Category.EditorIP
}

// LastModified function
func (c *CategorySchema) LastModified() *graphql.Time {
	return &graphql.Time{Time: c.Category.LastModified}
}

// IsDeleted function
func (c *CategorySchema) IsDeleted() *bool {
	return &c.Category.IsDeleted
}

// Deleted function
func (c *CategorySchema) Deleted() *graphql.Time {
	return &graphql.Time{Time: c.Category.Deleted}
}
