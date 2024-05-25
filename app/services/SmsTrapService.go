package services

import (
	"OmarFaruk-0x01/sms-trap/app/models"
	"OmarFaruk-0x01/sms-trap/app/utils"
	"context"
	"fmt"
	"sync"

	"github.com/uptrace/bun"
)

type SmsTrapService struct {
	db *bun.DB
}

func (sts *SmsTrapService) Create(trap *models.TrapDTO) error {
	newTrap := prepareTrap(trap)

	_, err := sts.db.NewInsert().Model(newTrap).Exec(context.Background())

	return err
}

func (sts *SmsTrapService) CreateMany(traps []*models.TrapDTO) error {

	wg := &sync.WaitGroup{}

	newTraps := make(chan *models.Trap, 1)
	preparedTraps := make([]*models.Trap, 0)

	wg.Add(len(traps))

	for _, trapData := range traps {
		go (func(wg *sync.WaitGroup, trapData *models.TrapDTO) {
			newTrap := prepareTrap(trapData)
			newTraps <- newTrap
			wg.Done()
		})(wg, trapData)
	}

	for i := 0; i < len(traps); i++ {
		preparedTraps = append(preparedTraps, <-newTraps)
	}

	wg.Wait()

	close(newTraps)

	fmt.Printf("prepared traps: %v\n: %v ", preparedTraps, len(traps))

	_, err := sts.db.NewInsert().Model(&preparedTraps).Exec(context.Background())

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

	err := sts.db.NewSelect().
		Table("traps").
		Where("phone = ?", phone).
		Limit(1).
		Scan(context.Background(), trap)

	if trap.Phone == "" {
		return nil, fmt.Errorf("not found")
	}

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

func (sts *SmsTrapService) Delete(trapId int64) error {
	res, err := sts.db.NewDelete().
		Model((*models.Trap)(nil)).
		Where("id = ?", trapId).
		Exec(context.Background())

	deletedCount, _ := res.RowsAffected()

	if deletedCount == 0 {
		return fmt.Errorf("not found")
	}

	return err
}

func (sts *SmsTrapService) DeleteAll() error {
	_, err := sts.db.NewDelete().
		Table("traps").
		Where("1 = 1").
		Exec(context.Background())

	return err
}

func NewSmsTrapService(db *bun.DB) *SmsTrapService {
	return &SmsTrapService{
		db,
	}
}

func prepareTrap(trap *models.TrapDTO) *models.Trap {
	return &models.Trap{
		Type:            utils.DetectEncoding(trap.Message),
		Phone:           trap.Phone,
		Label:           trap.Label,
		Message:         trap.Message,
		SMSCount:        utils.GetSMSCount(trap.Message),
		GSMCount:        utils.GetGSMCount(trap.Message),
		UnicodeCount:    utils.GetUnicodeCount(trap.Message),
		CharactersCount: utils.GetCharacterCount(trap.Message),
	}

}
