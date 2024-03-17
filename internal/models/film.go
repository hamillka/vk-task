package models

import (
	"time"
)

type Film struct {
	Id          int
	Name        string
	Description string
	ReleaseDate time.Time
	Rating      float32
	Actors      []*string
}

type FilmActors struct {
	filmId  int
	actorId int
}
