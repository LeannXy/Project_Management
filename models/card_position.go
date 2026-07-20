package models

import (
	"github.com/LeannXy/Project_Management/models/types"
	"github.com/google/uuid"
)

type CardPosition struct {
	InternalID int64           `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   uuid.UUID       `json:"public_id" db:"public_id" gorm:"type:uuid;not null"`
	ListID     int64           `json:"list_internal_id" db:"list_internal_id" gorm:"column:list_internal_id"`
	CardOrder  types.UUIDArray `json:"card_order" db:"card_order" gorm:"type:uuid[]"`
}
