package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" gorm:"column:owner_name;type:varchar(255); not null" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	BankID    string    `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number    string    `json:"number" gorm:"type:varchar(20); not null" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-" gorm:"ForeignKey:AccountID"`
}

func (a *Account) isValid() error {
	_, err := govalidator.ValidateStruct(a)

	if err != nil {
		return err
	}

	return nil
}

func NewAccount(bank *Bank, number, ownerName string) (*Account, error) {
	account := Account{
		OwnerName: ownerName,
		Number:    number,
		Bank:      bank,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()

	if err != nil {
		return nil, err
	}

	return &account, nil
}
