package repository

import (
	"fmt"
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/jmoiron/sqlx"
)

type clientUserRepository struct {
	db *sqlx.DB
}

func NewUserAndProfileRepository(db *sqlx.DB) domain.ClientUserRepo {
	return &clientUserRepository{
		db: db,
	}
}

func (r *clientUserRepository) GetUserWithProfile(clientUser *domain.ClientUser, userID string) error {
	if err := r.db.Get(clientUser,
		`select users.id, users.login, users.created_at, up.first_name, up.last_name, up.address, up.other_data 
	from users
    	left join user_profiles up on up.user_id = users.id
	where users.id = $1 limit 1`, userID); err != nil {
		return err
	}
	return nil
}

func (r *clientUserRepository) GetAuthClientUser(clientUser *domain.ClientUser, userID, token string) error {
	if err := r.db.Get(clientUser,
		`select users.id, users.login, users.created_at, a.token, up.first_name, up.last_name, up.address, up.other_data
	from users
	    left join access_tokens a on a.user_id = users.id
    	left join user_profiles up on up.user_id = users.id
	where users.id = $1 and a.token = $2 limit 1`, userID, token); err != nil {
		return err
	}
	return nil
}

func (r *clientUserRepository) GetAllWithOrderBy(field, orderBy string) (*sqlx.Rows, error) {
	query := fmt.Sprintf(`select
	users.id id,
	users.login as login,
	users.created_at as created_at,
	up.first_name as first_name,
	up.last_name as last_name,
	up.address as address,
	up.other_data
	from users
		left join user_profiles up on up.user_id = users.id order by %s %s`, field, orderBy)
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *clientUserRepository) GetAllWithOrderByAndFilter(filter map[string]string, field, orderBy string) (*sqlx.Rows, error) {
	query := `select 
    users.id id, 
    users.login as login, 
    users.created_at as created_at, 
    up.first_name as first_name, 
    up.last_name as last_name, 
    up.address as address,
	up.other_data 
	from users
    	left join user_profiles up on up.user_id = users.id where`

	i := 0
	for k, v := range filter {
		if i == 0 {
			if v == "notnull" || v == "isnull" {
				query = fmt.Sprintf("%s %s %s", query, k, v)
				continue
			}
			query = fmt.Sprintf("%s %s ilike '%s%s%s'", query, k, "%", v, "%")
		}
		if i > 0 {
			if v == "notnull" || v == "isnull" {
				query = fmt.Sprintf("%s and %s %s", query, k, v)
				continue
			}
			query = fmt.Sprintf("%s and %s ilike '%s%s%s'", query, k, "%", v, "%")
		}
		i++
	}

	query = fmt.Sprintf("%s order by %s %s", query, field, orderBy)
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
