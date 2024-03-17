package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
	"vk-task/internal/handlers/dto"
	"vk-task/internal/models"
)

type FilmsHandler struct {
	repository FilmRepository
}

func NewFilmsHandler(r FilmRepository) *FilmsHandler {
	return &FilmsHandler{
		repository: r,
	}
}

type FilmRepository interface {
	CreateFilm(name, description string, releaseDate time.Time,
		rating float32, filmActors []*dto.CreateOrUpdateActorRequestDto) (int, error)
}

func (h *FilmsHandler) CreateFilm(w http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	claims := ctx.Value("props").(jwt.MapClaims)
	role, _ := claims["role"].(int)
	if role == models.ADMIN {
		var filmReqDto dto.CreateOrUpdateFilmRequestDto
		err := json.NewDecoder(request.Body).Decode(&filmReqDto)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			errorDto := &dto.ErrorDto{
				Error: "Тело запроса некорректно",
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		releaseDate, err := time.Parse("2006-01-02", filmReqDto.ReleaseDate)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			errorDto := &dto.ErrorDto{
				Error: "Произошла ошибка",
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		_, err = h.repository.CreateFilm(filmReqDto.Name, filmReqDto.Description, releaseDate, filmReqDto.Rating, nil)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			errorDto := &dto.ErrorDto{
				Error: "Возникла внутренняя ошибка при запросе",
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		errorDto := &dto.ErrorDto{
			Error: "В доступе отказано",
		}
		err := json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}
}

func (h *FilmsHandler) UpdateFilm(w http.ResponseWriter, request *http.Request) {

}

func (h *FilmsHandler) DeleteFilm(w http.ResponseWriter, request *http.Request) {

}

func (h *FilmsHandler) GetFilms(w http.ResponseWriter, r *http.Request) {

}

func (h *FilmsHandler) GetFilmByName(w http.ResponseWriter, request *http.Request) {

}
