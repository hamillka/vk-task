package handlers

import (
	"net/http"
	"vk-task/internal/handlers/middlewares"
)

type Router struct {
	ar ActorRepository
	fr FilmRepository
	ur UserRepository
	ah *ActorsHandler
	fh *FilmsHandler
	uh *UsersHandler
}

func NewRouter(
	ar ActorRepository,
	fr FilmRepository,
	ur UserRepository,
) *Router {
	return &Router{
		ar: ar,
		fr: fr,
		ur: ur,
	}
}

func (r *Router) ProcessActorWithId(w http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPut:
		err := r.ah.UpdateActor()
		if err != nil {
			//w.Header().Add()
		}
	case http.MethodDelete:
		err := r.ah.DeleteActor()
		if err != nil {
			//w.Header().Add()
		}
	}
}

func (r *Router) ProcessActorWithoutId(w http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		r.ah.CreateActor(w, request)
	case http.MethodGet:
		r.ah.GetActorByName(w, request)
	}
}

func (r *Router) ProcessFilmWithoutId(w http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		r.fh.CreateFilm(w, request) // post   - addFilm
	case http.MethodGet:
		r.fh.GetFilmByName(w, request) // get - search film by name fragment
	}
}

func (r *Router) ProcessFilmRequestsWithId(w http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPut:
		r.fh.UpdateFilm(w, request)
	case http.MethodDelete:
		r.fh.DeleteFilm(w, request)
	}
}

func (r *Router) Route() {
	r.ah = NewActorsHandler(r.ar)

	http.Handle("/api/v1/actor", middlewares.AuthMiddleware(http.HandlerFunc(r.ProcessActorWithoutId)))
	http.Handle("/api/v1/actor/{id}", middlewares.AuthMiddleware(http.HandlerFunc(r.ProcessActorWithId))) // edit and delete actor

	http.Handle("/api/v1/actors", middlewares.AuthMiddleware(http.HandlerFunc(r.ah.GetAllActors))) // get all actors and list of films with them

	r.fh = NewFilmsHandler(r.fr)
	http.Handle("/api/v1/film", middlewares.AuthMiddleware(http.HandlerFunc(r.ProcessFilmWithoutId)))
	http.Handle("/api/v1/film/{id}", middlewares.AuthMiddleware(http.HandlerFunc(r.ProcessFilmRequestsWithId))) // edit and delete actor

	http.Handle("/api/v1/films", middlewares.AuthMiddleware(http.HandlerFunc(r.fh.GetFilms))) // get    - getFilms

	//http.HandleFunc("/api/v1/films?orderBy=name", tmp)        // get    - getFilmsOrderedByName
	//http.HandleFunc("/api/v1/films?orderBy=rating", tmp)      // get    - getFilmsOrderedByRating
	//http.HandleFunc("/api/v1/films?orderBy=releasedate", tmp) // get    - getFilmsOrderedByReleaseDate

	r.uh = NewUsersHandler(r.ur)
	http.HandleFunc("/api/v1/login", r.uh.LoginUser) // loginMethod
}
