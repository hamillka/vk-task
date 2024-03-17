package dto

var JwtSecretKey = []byte("secret-key")

type LoginResponseDto struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
