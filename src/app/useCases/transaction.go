package useCases

import (
	"github.com/iago-f-s-e/pix-code-go/src/domain/model"
)

type TransactionUseCases struct {
	TransactionRepository model.TransactionRepository
	PixKeyRepository      model.PixKeyRepository
}

func (p *TransactionUseCases) Register(accountId, pixKeyTo, pixKeyKindTo, description string, amount float64) (*model.Transaction, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := p.PixKeyRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, pixKey, amount, description)
	if err != nil {
		return nil, err
	}

	err = p.TransactionRepository.Register(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (p *TransactionUseCases) Confirm(id string) error {
	transaction, err := p.TransactionRepository.Find(id)

	if err != nil {
		return err
	}

	transaction.Confirm()

	return p.TransactionRepository.Save(transaction)
}

func (p *TransactionUseCases) Complete(id string) error {
	transaction, err := p.TransactionRepository.Find(id)

	if err != nil {
		return err
	}

	transaction.Complete()

	return p.TransactionRepository.Save(transaction)
}

func (p *TransactionUseCases) Cancel(id, description string) error {
	transaction, err := p.TransactionRepository.Find(id)

	if err != nil {
		return err
	}

	transaction.Cancel(description)

	return p.TransactionRepository.Save(transaction)
}

func (p *TransactionUseCases) Error(id, reason string) (*model.Transaction, error) {
	transaction, err := p.TransactionRepository.Find(id)

	if err != nil {
		return nil, err
	}

	transaction.Error(reason)

	err = p.TransactionRepository.Save(transaction)

	return transaction, err
}

func (p *TransactionUseCases) Find(id string) (*model.Transaction, error) {
	return p.TransactionRepository.Find(id)
}
