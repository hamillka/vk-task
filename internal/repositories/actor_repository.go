package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"vk-task/internal/handlers/dto"
	"vk-task/internal/models"
)

type PostgresActorRepository struct {
	db *sqlx.DB
}

const (
	findActor   = "SELECT * FROM actors WHERE name LIKE $1"
	createActor = "INSERT INTO actors (name, sex, birthdate) VALUES ($1, $2, $3) RETURNING id"
)

func (r *PostgresActorRepository) FindActor(nameFragment string) (*dto.GetActorsByNameFragmentResponseDto, error) {
	var actors []*models.Actor
	rows, err := r.db.Query(findActor, "%"+nameFragment+"%")
	if err != nil {
		return nil, ErrRecordNotFound
	}

	if err := rows.Err(); err != nil {
		return nil, ErrDatabaseReadingError
	}

	for rows.Next() {
		actor := new(models.Actor)
		if err := rows.Scan(&actor.Id, &actor.Name, &actor.Sex, &actor.BirthDate); err != nil {
			return nil, ErrDatabaseReadingError
		}
		actors = append(actors, actor)
	}
	fmt.Println(actors)
	defer rows.Close()
	return dto.ConvertActorToDto(actors), nil
}

func (r *PostgresActorRepository) CreateActor(name, sex string, birthDate time.Time) (int, error) {
	var id int

	row := r.db.QueryRow(createActor, name, sex, birthDate)
	if err := row.Scan(&id); err != nil {
		return 0, ErrDatabaseWritingError
	}

	return id, nil
}

func NewActorRepository(db *sqlx.DB) *PostgresActorRepository {
	return &PostgresActorRepository{
		db: db,
	}
}
