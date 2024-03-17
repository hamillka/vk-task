APP=filmoteka

build:
	docker-compose build $(APP)

run:
	docker-compose up $(APP)

clean:
	docker-compose down
