package repositories

import (
	"fmt"

	"github.com/iago-f-s-e/full-cycle-go/src/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (r *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	return r.Db.Create(transaction).Error
}

func (r *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	return r.Db.Save(transaction).Error
}

func (r *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("transaction not found")
	}

	return &transaction, nil
}
