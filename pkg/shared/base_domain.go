package shared

import (
	"time"
)

// BaseDomain structure
type BaseDomain struct {
	CreatorID    string    `gorm:"column:CREATOR_ID"`
	CreatorIP    string    `gorm:"column:CREATOR_IP"`
	Created      time.Time `gorm:"column:CREATED"`
	EditorID     string    `gorm:"column:EDITOR_ID"`
	EditorIP     string    `gorm:"column:EDITOR_IP"`
	LastModified time.Time `gorm:"column:LAST_MODIFIED"`
	IsDeleted    bool      `gorm:"column:IS_DELETED"`
	Deleted      time.Time `gorm:"column:DELETED"`
}
