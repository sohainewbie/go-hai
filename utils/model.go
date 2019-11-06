package utils

import "time"

type BaseModel struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}

type BaseSourceByModel struct {
	CreatedBy uint64 `json:"createdBy,omitempty"`
	ChangedBy uint64 `json:"changedBy,omitempty"`
	DeletedBy uint64 `json:"deletedBy,omitempty"`
}
