package models

import "time"

type Actor struct {
	Id        int
	Name      string
	Sex       string
	BirthDate time.Time
}
