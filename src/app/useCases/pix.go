package useCases

import (
	"github.com/iago-f-s-e/pix-code-go/src/domain/model"
	"github.com/iago-f-s-e/pix-code-go/src/infra/repositories"
)

type PixUseCases struct {
	PixKeyRepository repositories.PixKeyRepositoryDb
}

func (p *PixUseCases) RegisterKey(key, kind, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(account, kind, key)

	if err != nil {
		return nil, err
	}

	return p.PixKeyRepository.Register(pixKey)
}

func (p *PixUseCases) FindKey(key, kind string) (*model.PixKey, error) {
	return p.PixKeyRepository.FindKeyByKind(key, kind)
}
