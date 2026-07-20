package models

type CardAssignee struct{
	CardID int64 `json:"card_internal_id" db:"card_internal_id" gorm:"card_internal_id"`
	UserID int64 `json:"user_internal_id" db:"user_internal_id" gor:"user_internal_id" `
}