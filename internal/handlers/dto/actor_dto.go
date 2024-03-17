package dto

import (
	"time"
	"vk-task/internal/models"
)

type ActorDto struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	BirthDate string `json:"birthdate"`
}

type CreateOrUpdateActorRequestDto struct {
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	BirthDate string `json:"birthdate"`
}

type GetActorByNameFragmentRequestDto struct {
	NameFragment string `json:"name"`
}

type GetActorByNameFragmentDto struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	BirthDate string `json:"birthdate"`
}

type GetActorsByNameFragmentResponseDto struct {
	Actors []*GetActorByNameFragmentDto `json:"actors"`
}

func ConvertActorToDto(actors []*models.Actor) *GetActorsByNameFragmentResponseDto {
	var actorsDto []*GetActorByNameFragmentDto

	for _, val := range actors {
		actorByNameFragment := &GetActorByNameFragmentDto{
			Id:        val.Id,
			Name:      val.Name,
			Sex:       val.Sex,
			BirthDate: val.BirthDate.String(),
		}
		actorsDto = append(actorsDto, actorByNameFragment)
	}

	if actors == nil {
		actorsDto = []*GetActorByNameFragmentDto{}
	}

	return &GetActorsByNameFragmentResponseDto{
		Actors: actorsDto,
	}
}

func ConvertDtoToActor(actorsDto []*CreateOrUpdateActorRequestDto) (actors []*models.Actor) {
	for _, val := range actorsDto {
		t, _ := time.Parse("2006-01-02", val.BirthDate)
		actor := &models.Actor{
			Name:      val.Name,
			Sex:       val.Sex,
			BirthDate: t,
		}

		actors = append(actors, actor)
	}

	if actorsDto == nil {
		actors = []*models.Actor{}
	}

	return
}
