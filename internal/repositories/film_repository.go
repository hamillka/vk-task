package repositories

import (
	"github.com/jmoiron/sqlx"
	"time"
	"vk-task/internal/handlers/dto"
	"vk-task/internal/models"
)

type PostgresFilmRepository struct {
	db *sqlx.DB
}

func NewFilmRepository(db *sqlx.DB) *PostgresFilmRepository {
	return &PostgresFilmRepository{
		db: db,
	}
}

const (
	findFilm   = "SELECT * FROM films WHERE name LIKE $1"
	createFilm = "INSERT INTO films (name, description, releasedate, rating) VALUES ($1, $2, $3, $4) RETURNING id"
)

func (r *PostgresFilmRepository) FindFilm(nameFragment string) (*dto.GetFilmsByNameFragmentResponseDto, error) {
	var films []*models.Film
	rows, err := r.db.Query(findFilm, "%"+nameFragment+"%")
	if err != nil {
		return nil, ErrRecordNotFound
	}

	if err := rows.Err(); err != nil {
		return nil, ErrDatabaseReadingError
	}

	for rows.Next() {
		film := new(models.Film)
		if err := rows.Scan(); err != nil {
			return nil, ErrDatabaseReadingError
		}
		films = append(films, film)
	}
	defer rows.Close()
	return dto.ConvertFilmToDto(films), nil
}

func (r *PostgresFilmRepository) CreateFilm(
	name, description string,
	releaseDate time.Time,
	rating float32,
	filmActors []*dto.CreateOrUpdateActorRequestDto,
) (int, error) {
	//var actors []*models.Actor

	var id int

	row := r.db.QueryRow(createFilm, name, description, releaseDate, rating)
	if err := row.Scan(&id); err != nil {
		return 0, ErrDatabaseWritingError
	}

	return id, nil
}
