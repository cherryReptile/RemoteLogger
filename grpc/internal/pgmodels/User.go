package pgmodels

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	BaseModel
	ID        string    `json:"id" db:"id"`
	Login     string    `json:"login" db:"login"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *User) Create(db *sqlx.DB, provider string) error {
	u.ID = uuid.NewString()
	u.CreatedAt = time.Now()

	_, err := db.NamedExec(`INSERT INTO users (id, login, password, created_at) 
								VALUES (:id, :login, :password, :created_at)`, u)

	if err != nil {
		//return errors.New("failed to create user " + err.Error())
		return err
	}

	// update model
	if err = db.Get(u, "SELECT * FROM users WHERE id=$1 LIMIT 1", u.ID); err != nil {
		return err
	}

	authProvider := new(AuthProvider)
	if err = authProvider.GetByProvider(db, provider); err != nil {
		return err
	}

	intermediate := new(Intermediate)
	if err = intermediate.Create(db, u.ID, authProvider.ID); err != nil {
		return err
	}

	return nil
}

func (u *User) FindByUUID(db *sqlx.DB, uuid string) error {
	if err := db.Get(u, "SELECT * FROM users WHERE id=$1", uuid); err != nil {
		return err
	}
	return nil
}

func (u *User) CheckOnExistsWithoutPassword(db *sqlx.DB, login, provider string) error {
	if err := db.Get(u,
		`select users.id, login, password, users.created_at 
	from users
    	left join intermediate i on i.user_id = users.id
    	left join auth_providers ap on i.provider_id = ap.id
	where users.login = $1 and ap.provider = $2 limit 1`, login, provider); err != nil {
		return err
	}
	return nil
}

//func (u *User) FindByUniqueAndService(db *sqlx.DB, unique, service string) error {
//	if err := db.Get(u, "SELECT * FROM users WHERE unique_raw=$1 AND authorized_by=$2", unique, service); err != nil {
//		return err
//	}
//
//	return nil
//}

func (u *User) GetTokenByStr(db *sqlx.DB, token string) (*AccessToken, error) {
	t := AccessToken{}
	err := db.Get(&t, "SELECT * FROM access_tokens WHERE user_id=$1 AND token=$2 ORDER BY id DESC", u.ID, token)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
