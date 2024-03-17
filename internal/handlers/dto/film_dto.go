package dto

import "vk-task/internal/models"

type FilmDto struct {
	Id          int                              `json:"id"`
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	ReleaseDate string                           `json:"releaseDate"`
	Rating      float32                          `json:"rating"`
	Actors      []*CreateOrUpdateActorRequestDto `json:"actors"` //// ???????????
}

type CreateOrUpdateFilmRequestDto struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ReleaseDate string    `json:"releaseDate"`
	Rating      float32   `json:"rating"`
	Actors      []*string `json:"actors"`
}

type GetFilmByNameFragmentDto struct {
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	ReleaseDate string                           `json:"releaseDate"`
	Rating      float32                          `json:"rating"`
	Actors      []*CreateOrUpdateActorRequestDto `json:"actors"`
}

type GetFilmsByNameFragmentResponseDto struct {
	Films []*GetFilmByNameFragmentDto `json:"films"`
}

func ConvertFilmToDto(films []*models.Film) *GetFilmsByNameFragmentResponseDto {
	var filmsDto []*GetFilmByNameFragmentDto

	for _, val := range films {
		filmByNameFragment := &GetFilmByNameFragmentDto{
			Name:        val.Name,
			Description: val.Description,
			ReleaseDate: val.ReleaseDate.String(),
			Rating:      val.Rating,
			Actors:      nil,
		}
		filmsDto = append(filmsDto, filmByNameFragment)
	}

	if films == nil {
		filmsDto = []*GetFilmByNameFragmentDto{}
	}

	return &GetFilmsByNameFragmentResponseDto{
		Films: filmsDto,
	}

}
