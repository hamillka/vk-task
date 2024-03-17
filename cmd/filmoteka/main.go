package main

import (
	"net/http"
	"vk-task/internal/config"
	"vk-task/internal/db"
	"vk-task/internal/handlers"
	"vk-task/internal/logger"
	"vk-task/internal/repositories"
)

func main() {
	config, err := config.New()
	logger := logger.CreateLogger(config.Log)

	//defer func() {
	//	err := logger.Sync()
	//	if err != nil {
	//		logger.Errorf("Error while syncing logger: %v", err)
	//	}
	//}()

	if err != nil {
		logger.Errorf("Something went wrong with config: %v", err)
	}

	db, err := db.CreateConnection(config.Db)

	defer func() {
		err := db.Close()
		if err != nil {
			logger.Errorf("Error while closing connection to db: %v", err)
		}
	}()

	if err != nil {
		logger.Fatalf("Error while connecting to database: %v", err)
	}

	ar := repositories.NewActorRepository(db)
	ur := repositories.NewUserRepository(db)
	fr := repositories.NewFilmRepository(db)
	r := handlers.NewRouter(ar, fr, ur)
	r.Route()

	port := config.Port
	logger.Info("Server is started on port ", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Fatalf("Error while starting server: %v", err)
	}
}
