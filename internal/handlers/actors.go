package handlers

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
	"vk-task/internal/handlers/dto"
	"vk-task/internal/models"
	"vk-task/internal/repositories"
)

type ActorsHandler struct {
	repository ActorRepository
}

func NewActorsHandler(r ActorRepository) *ActorsHandler {
	return &ActorsHandler{
		repository: r,
	}
}

type ActorRepository interface {
	FindActor(nameFragment string) (*dto.GetActorsByNameFragmentResponseDto, error)
	CreateActor(name, sex string, birthDate time.Time) (int, error)
}

func (h *ActorsHandler) CreateActor(w http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	claims := ctx.Value("props").(jwt.MapClaims)
	role, _ := claims["role"].(int)
	if role == models.ADMIN {
		var actorDto dto.CreateOrUpdateActorRequestDto
		err := json.NewDecoder(request.Body).Decode(&actorDto)
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
		birthDate, err := time.Parse("2006-01-02", actorDto.BirthDate)
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
		_, err = h.repository.CreateActor(actorDto.Name, actorDto.Sex, birthDate) // post   - addActor
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

func (h *ActorsHandler) UpdateActor() error {

	return nil
}

func (h *ActorsHandler) DeleteActor() error {

	return nil
}

func (h *ActorsHandler) GetActorByName(w http.ResponseWriter, request *http.Request) {
	var actorDto dto.GetActorByNameFragmentRequestDto
	err := json.NewDecoder(request.Body).Decode(&actorDto)
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
	actors, err := h.repository.FindActor(actorDto.NameFragment)
	if err != nil {
		switch {
		case errors.Is(err, repositories.ErrRecordNotFound):
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			errorDto := &dto.ErrorDto{
				Error: "Запись с таким именем в таблице актеров не найдена",
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			errorDto := &dto.ErrorDto{
				Error: "Возникла внутренняя ошибка при запросе актера по фрагементу имени",
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(actors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *ActorsHandler) GetAllActors(writer http.ResponseWriter, request *http.Request) {
	//w.Header().Add()
}
