package repositories

import (
	"fmt"

	"github.com/iago-f-s-e/pix-code-go/src/domain/model"
	"gorm.io/gorm"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (r *PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	return r.Db.Create(bank).Error
}

func (r *PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	return r.Db.Create(account).Error
}

func (r *PixKeyRepositoryDb) Register(pixkey *model.PixKey) (*model.PixKey, error) {
	return pixkey, r.Db.Create(pixkey).Error
}

func (r *PixKeyRepositoryDb) FindKeyByKind(key, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(&pixKey, "key = ? and kind = ?", key, kind)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("pix key not found")
	}

	return &pixKey, nil
}

func (r *PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	r.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("account not found")
	}

	return &account, nil
}

func (r *PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	r.Db.Preload("Bank").First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("bank not found")
	}

	return &bank, nil
}
