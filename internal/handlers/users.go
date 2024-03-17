package handlers

import (
	"encoding/json"
	goErrors "errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
	"vk-task/internal/handlers/dto"
	"vk-task/internal/models"
	"vk-task/internal/repositories"
)

type UsersHandler struct {
	repository UserRepository
}

func NewUsersHandler(r UserRepository) *UsersHandler {
	return &UsersHandler{
		repository: r,
	}
}

type UserRepository interface {
	GetUserByLoginAndPassword(login, password string) (*models.User, error)
}

func createToken(role int) (string, error) {
	payload := jwt.MapClaims{
		"role": role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString(dto.JwtSecretKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (h *UsersHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginReq dto.LoginRequest

	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные входные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	user, err := h.repository.GetUserByLoginAndPassword(loginReq.Login, loginReq.Password)
	if err != nil {
		switch {
		case goErrors.Is(err, repositories.ErrRecordNotFound):
			w.WriteHeader(http.StatusUnauthorized)
			errorDto := &dto.ErrorDto{
				Error: "Неверный логин и/или пароль",
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			errorDto := &dto.ErrorDto{
				Error: "Возникла внутренняя ошибка при запросе",
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

		return
	}

	t, err := createToken(user.Role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorDto := &dto.ErrorDto{
			Error: "Возникла внутренняя ошибка при генерации токена",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	LoginResponseDto := dto.LoginResponseDto{
		Token: t,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(LoginResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
