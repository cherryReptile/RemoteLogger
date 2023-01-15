package usecase

import "github.com/cherryReptile/WS-AUTH/domain"

type providerUsecase struct {
	providerRepo domain.ProviderRepo
}

func NewProviderUsecase(pr domain.ProviderRepo) domain.ProviderUsecase {
	return &providerUsecase{
		providerRepo: pr,
	}
}

func (u *providerUsecase) GetByProvider(provider *domain.Provider, name string) error {
	return u.providerRepo.GetByProvider(provider, name)
}
