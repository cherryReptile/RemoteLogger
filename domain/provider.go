package domain

type Provider struct {
	ID        uint   `db:"id"`
	Provider  string `db:"provider"`
	UniqueKey string `db:"unique_key"`
}

type ProviderRepo interface {
	GetByProvider(provider *Provider, name string) error
}

type ProviderUsecase interface {
	GetByProvider(provider *Provider, name string) error
}
