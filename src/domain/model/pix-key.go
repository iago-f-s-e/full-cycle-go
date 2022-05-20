package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKeyRepository interface {
	Register(pixKey *PixKey) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindKeyByKind(key, kind string) (*PixKey, error)
	FindAccount(id string) (*Account, error)
	FindBank(id string) (*Bank, error)
}

type PixKey struct {
	Base      `valid:"required"`
	King      string   `json:"king" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (p *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(p)

	isValidKind := p.King == "email" || p.King == "cpf"
	isValidStatus := p.Status == "active" || p.Status == "inactive"

	if !isValidKind {
		return errors.New("invalid type of key")
	}

	if !isValidStatus {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewPixKey(account *Account, kind, key string) (*PixKey, error) {
	pixKey := PixKey{
		Key:     key,
		King:    kind,
		Account: account,
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()
	pixKey.Status = "active"

	err := account.isValid()

	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
