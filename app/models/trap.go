package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Trap struct {
	bun.BaseModel `bun:"table:traps,alias:t"`

	Id      int64  `bun:"id,pk,autoincrement"`
	Phone   string `bun:"phone,notnull,nullzero,"`
	Message string `bun:"message,notnull,nullzero"`
	Type    string `bun:"type"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type TrapDTO struct {
	Phone   string
	Message string
	Type    string
}

type TrapPhones struct {
	Phone     string    `json:"phone"`
	Count     string    `json:"count"`
	LastSmsAt time.Time `json:"last_sms_at"`
}
