SELECT 'CREATE DATABASE "filmoteka"'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'filmoteka')\gexec

\connect "filmoteka";

CREATE TABLE IF NOT EXISTS actors (
    id serial PRIMARY KEY,
    name text NOT NULL,
    sex text NOT NULL,
    birthdate date NOT NULL
);

CREATE TABLE IF NOT EXISTS films (
    id serial PRIMARY KEY,
    name text CHECK ( char_length(name) between 1 and 150 ) NOT NULL,
    description varchar(1000),
    releasedate date NOT NULL,
    rating float4 CHECK ( rating between 0 and 10 ) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    role integer CHECK ( role between 0 and 1 ) NOT NULL,
    login text UNIQUE NOT NULL,
    password text NOT NULL
);

CREATE TABLE IF NOT EXISTS films_actors (
    film_id integer NOT NULL,
    actor_id integer NOT NULL,
    CONSTRAINT fk_film FOREIGN KEY (film_id) REFERENCES films (id),
    CONSTRAINT fk_actor FOREIGN KEY (actor_id) REFERENCES actors (id)
);