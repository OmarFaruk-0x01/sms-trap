package migration

import (
	"OmarFaruk-0x01/sms-trap/app/models"
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

type BunMigration struct {
	db *bun.DB
}

func (m *BunMigration) Up() {

	_, err := m.db.NewCreateTable().
		Model((*models.Trap)(nil)).
		IfNotExists().
		Exec(context.Background())

	if err != nil {
		panic(err)
	}

	fmt.Println("Created table: traps")

}

func (m *BunMigration) Down() {

	_, err := m.db.NewDropTable().
		Model((*models.Trap)(nil)).
		IfExists().
		Exec(context.Background())

	if err != nil {
		panic(err)
	}

	fmt.Println("Dropped table: traps")

}

func NewBunMigration(db *bun.DB) Migration {
	return &BunMigration{db}
}
