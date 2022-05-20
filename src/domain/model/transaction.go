package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   = "pending"
	TransactionCompleted = "completed"
	TransactionError     = "error"
	TransactionConfirmed = "confirmed"
)

type TransactionRepository interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base              `valid:"-"`
	AccountFrom       *Account `valid:"-" gorm:"column:account_from_id; type:uuid"`
	Amount            float64  `json:"amount" gorm:"type:float; not null" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyIdTo        string   `gorm:"column:pix_key_id_to;type:uuid;not null" valid:"-"`
	Status            string   `json:"status" gorm:"type:varchar(20); not null" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255); not null" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" gorm:"column:cancel_description;type:varchar(255); not null" valid:"notnull"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	isValidAmount := t.Amount > 0
	isValidStatus := t.Status == TransactionPending || t.Status == TransactionCompleted || t.Status == TransactionError || t.Status == TransactionConfirmed
	isValidTransaction := t.PixKeyTo.ID != t.AccountFrom.ID

	if !isValidAmount {
		return errors.New("the amount must be greater than 0")
	}

	if !isValidStatus {
		return errors.New("invalid status for the transaction")
	}

	if !isValidTransaction {
		return errors.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()

	return t.isValid()
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()

	return t.isValid()
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.CancelDescription = description
	t.UpdatedAt = time.Now()

	return t.isValid()
}

func NewTransaction(accountFrom *Account, pixKeyTo *PixKey, amount float64, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountFrom,
		PixKeyTo:    pixKeyTo,
		Amount:      amount,
		Description: description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()
	transaction.Status = TransactionPending

	err := transaction.isValid()

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
