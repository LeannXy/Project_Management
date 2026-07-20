package models

import (
	"github.com/LeannXy/Project_Management/models/types"
	"github.com/google/uuid"
)

type ListPosition struct {
	InternalID int64 `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" db:"public_id" `
	BoardID int64 `json:"board_internal_id" db:"board_internal_id" gorm:"column:board_internal_id"`
	//buat type data
	ListOrder types.UUIDArray `json:"list_order"`
}