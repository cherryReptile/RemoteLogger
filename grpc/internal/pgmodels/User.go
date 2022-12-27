package pgmodels

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	BaseModel
	ID        string    `json:"id" db:"id"`
	Login     string    `json:"login" db:"login"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *User) Create(db *sqlx.DB, providerID uint) error {
	u.ID = uuid.NewString()
	u.CreatedAt = time.Now()

	_, err := db.NamedExec(`INSERT INTO users (id, login, created_at) 
								VALUES (:id, :login, :created_at)`, u)

	if err != nil {
		//return errors.New("failed to create user " + err.Error())
		return err
	}

	// update model
	if err = db.Get(u, "SELECT * FROM users WHERE id=$1 LIMIT 1", u.ID); err != nil {
		return err
	}

	intermediate := new(Intermediate)
	if err = intermediate.Create(db, u.ID, providerID); err != nil {
		return err
	}

	return nil
}

func (u *User) Find(db *sqlx.DB, uuid string) error {
	if err := db.Get(u, "SELECT * FROM users WHERE id=$1", uuid); err != nil {
		return err
	}
	return nil
}

func (u *User) FindByLoginAndProvider(db *sqlx.DB, login, provider string) error {
	if err := db.Get(u,
		`select users.id, login, users.created_at 
	from users
    	left join users_providers up on up.user_id = users.id
    	left join providers p on up.provider_id = p.id
	where users.login = $1 and p.provider = $2 limit 1`, login, provider); err != nil {
		return err
	}
	return nil
}

func (u *User) GetProviderData(db *sqlx.DB, provider string) (*ProvidersData, error) {
	ap := new(Provider)
	pd := new(ProvidersData)
	ap.GetByProvider(db, provider)

	if ap.ID == 0 {
		return nil, errors.New("unknown provider")
	}

	pd.FindByUserUUIDAndProviderID(db, u.ID, pd.ID)
	return pd, nil
}

func (u *User) GetTokenByStr(db *sqlx.DB, token string) (*AccessToken, error) {
	t := AccessToken{}
	err := db.Get(&t, "SELECT * FROM access_tokens WHERE user_id=$1 AND token=$2 ORDER BY id DESC", u.ID, token)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
