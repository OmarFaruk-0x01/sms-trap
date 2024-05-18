package services

import (
	"OmarFaruk-0x01/sms-trap/app/models"
	"context"

	"github.com/uptrace/bun"
)

type SmsTrapService struct {
	db *bun.DB
}

func (sts *SmsTrapService) Create(trap *models.TrapDTO) error {
	newTrap := &models.Trap{
		Phone:   trap.Phone,
		Message: trap.Message,
		Type:    trap.Type,
	}
	_, err := sts.db.NewInsert().Model(newTrap).Exec(context.Background())

	return err
}

func (sts *SmsTrapService) CreateMany(traps []*models.TrapDTO) error {

	newTraps := []*models.Trap{}

	for _, trapData := range traps {
		newTraps = append(newTraps, &models.Trap{
			Phone:   trapData.Phone,
			Message: trapData.Message,
			Type:    trapData.Type,
		})
	}
	_, err := sts.db.NewInsert().Model(&newTraps).Exec(context.Background())

	return err
}

func (sts *SmsTrapService) Find(trapId int64) (*models.Trap, error) {

	trap := &models.Trap{}

	_, err := sts.db.NewSelect().
		Model(trap).
		Where("id = ?", trapId).
		Limit(1).
		Exec(context.Background())

	return trap, err
}

func (sts *SmsTrapService) FindByPhone(phone string) (*models.Trap, error) {

	trap := &models.Trap{}

	_, err := sts.db.NewSelect().
		Model(trap).
		Where("phone = ?", phone).
		Limit(1).
		Exec(context.Background())

	return trap, err
}

func (sts *SmsTrapService) FindAll() ([]*models.Trap, error) {

	traps := make([]*models.Trap, 0)

	err := sts.db.NewSelect().
		Model(&traps).
		Scan(context.Background())

	return traps, err
}

func (sts *SmsTrapService) FindAllByPhone(phone string) ([]*models.Trap, error) {

	traps := make([]*models.Trap, 0)

	err := sts.db.NewSelect().
		Model(&traps).
		Where("phone = ?", phone).
		Order("created_at DESC").
		Scan(context.Background())

	return traps, err
}

func (sts *SmsTrapService) GetPhones() ([]*models.TrapPhones, error) {

	traps := make([]*models.TrapPhones, 0)

	err := sts.db.NewSelect().
		TableExpr("traps").
		ColumnExpr("phone").
		ColumnExpr("COUNT(id) as count").
		ColumnExpr("MAX(created_at) as last_sms_at").
		Group("phone").
		Scan(context.Background(), &traps)

	return traps, err
}

func NewSmsTrapService(db *bun.DB) *SmsTrapService {
	return &SmsTrapService{
		db,
	}
}
