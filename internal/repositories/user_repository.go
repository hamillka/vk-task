package repositories

import (
	goErrors "errors"

	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"vk-task/internal/models"
)

type PostgresUserRepository struct {
	db *sqlx.DB
}

const (
	selectUserByLoginAndPassword = `SELECT id, role, login, password FROM users WHERE login = $1 AND password = $2;`
)

var (
	ErrRecordNotFound       = goErrors.New("Record was not found")
	ErrDatabaseWritingError = goErrors.New("Error while writing to DB")
	ErrDatabaseReadingError = goErrors.New("Error while reading from DB")
	ErrRecordAlreadyExists  = goErrors.New("Record with this data already exists")
)

func (r *PostgresUserRepository) GetUserByLoginAndPassword(login, password string) (*models.User, error) {
	user := new(models.User)
	err := r.db.QueryRow(selectUserByLoginAndPassword, login, password).Scan(&user.Id, &user.Role, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
	}

	return user, nil
}

func NewUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Login() error {
	return nil
}
