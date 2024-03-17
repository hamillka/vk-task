Актеры:
-
1. добавление информации об актере (имя, пол, дата рождения) - POST
2. изменение информации об актере (как частично, так и полностью) - PUT
3. удаление информации об актере - DELETE

Фильмы:
-
1. добавление информации о фильме (название, описание, дата выпуска, рейтинг, список актеров) - POST
2. изменение информации о фильме (как частично, так и полностью) - PUT
3. удаление информации о фильме - DELETE
4. получить список фильмов с возможностью сортировки по названию/рейтингу/дате выпуска. По умолчанию - по рейтингу (убыв)
5. поиск фильма по фрагменту названия, по фрагменту имени актера ????
6. получить список актеров, для каждого актера выдается список фильмов с его участием
7. авторизация
8. две роли - юзер и админ. Юзер - получение данных и поиск, админ - все действия.

/login (в теле запроса отправляются пароль и логин) -> handlerLogin: (вызов метода userRepository, который проверяет, что есть юзер с таким логином и паролем, возвращает user struct (4 поля)), в самом LoginHandlere генерация jwt токена, в котором хранится роль пользователя, в w возвращается токен.
Этот токен вставляется в header по ключу (x-auth) -> 1) middleware 2) распарсить токен и чек роли в каждом хендлере

Структуры:
- 
actors:
{
	id
	name
	sex
	birthdate
}

films:
{
	id
	name
	description
	releasedate
	rating
}

filmsActors:
{
	filmId
	actorId
}

users:
{
	id
	login
	password
	role
}