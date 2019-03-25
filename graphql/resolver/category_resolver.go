package resolver

import (
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/musobarlab/gorengan/modules/product/domain"
)

// CategoryInput struct
type CategoryInput struct {
	ID           string
	Name         string
	CreatorID    string
	CreatorIP    string
	Created      *time.Time
	EditorID     string
	EditorIP     string
	LastModified *time.Time
	IsDeleted    *bool
	Deleted      *time.Time
}

// CategoryResolver resolver
type CategoryResolver struct {
	c *domain.Category
}

// ID function
func (c *CategoryResolver) ID() graphql.ID {
	return graphql.ID(c.c.ID)
}

// Name function
func (c *CategoryResolver) Name() string {
	return c.c.Name
}

// CreatorID function
func (c *CategoryResolver) CreatorID() string {
	return c.c.CreatorID
}

// CreatorIP function
func (c *CategoryResolver) CreatorIP() string {
	return c.c.CreatorIP
}

// Created function
func (c *CategoryResolver) Created() *graphql.Time {
	return &graphql.Time{Time: c.c.Created}
}

// EditorID function
func (c *CategoryResolver) EditorID() string {
	return c.c.EditorID
}

// EditorIP function
func (c *CategoryResolver) EditorIP() string {
	return c.c.EditorIP
}

// LastModified function
func (c *CategoryResolver) LastModified() *graphql.Time {
	return &graphql.Time{Time: c.c.LastModified}
}

// IsDeleted function
func (c *CategoryResolver) IsDeleted() *bool {
	return &c.c.IsDeleted
}

// Deleted function
func (c *CategoryResolver) Deleted() *graphql.Time {
	return &graphql.Time{Time: c.c.Deleted}
}
