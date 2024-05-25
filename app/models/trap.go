package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Trap struct {
	bun.BaseModel `json:"-" bun:"table:traps,alias:t"`

	Id      int64  `json:"id" bun:"id,pk,autoincrement"`
	Phone   string `json:"phone" bun:"phone,notnull,nullzero,"`
	Message string `json:"message" bun:"message,notnull,nullzero"`
	Type    string `json:"type" bun:"type"`
	Label   string `json:"label" bun:"label"`

	CharactersCount int `json:"characters_count" bun:"characters_count"`
	SMSCount        int `json:"sms_count" bun:"sms_count"`
	GSMCount        int `json:"gsm_count" bun:"gsm_count"`
	UnicodeCount    int `json:"unicode_count" bun:"unicode_count"`

	CreatedAt time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
}

type TrapDTO struct {
	Phone   string
	Message string
	Label   string
}

type TrapPhones struct {
	Phone     string    `json:"phone"`
	Count     string    `json:"count"`
	LastSmsAt time.Time `json:"last_sms_at"`
}

var TrapTypeColor = map[string]string{
	"text":    "primary",
	"unicode": "danger",
}

var TrapLabelColor = map[string]string{
	"transactional": "warning",
	"promotional":   "primary",
}
