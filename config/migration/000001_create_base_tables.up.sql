DROP TYPE IF EXISTS UNIVERSE;
CREATE TYPE UNIVERSE AS ENUM ('DC', 'MARVEL', 'DC|MARVEL');

CREATE TABLE team(
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    universe UNIVERSE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE character(
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    character_name VARCHAR(100) UNIQUE NOT NULL,
    civil_name VARCHAR(100),
    hero BOOLEAN,
    universe UNIVERSE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    fk_team UUID,
    FOREIGN KEY (fk_team) REFERENCES team(id)
);


CREATE TABLE ability(
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    description VARCHAR(300) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE game(
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    release_year SMALLINT NOT NULL,
    universe UNIVERSE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE hq(
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    release_year SMALLINT NOT NULL,
    universe UNIVERSE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE serie(
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    release_year SMALLINT NOT NULL,
    season SMALLINT NOT NULL,
    chapters SMALLINT NOT NULL,
    universe UNIVERSE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE movie (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    release_year SMALLINT NOT NULL,
    universe UNIVERSE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE console (
    name VARCHAR(20) PRIMARY KEY
);

CREATE TABLE character_serie(
    fk_character UUID,
    fk_serie UUID,
    CONSTRAINT pk_character_serie PRIMARY KEY (fk_character,fk_serie),
    FOREIGN KEY (fk_character) REFERENCES character(id),
    FOREIGN KEY (fk_serie) REFERENCES serie(id)
);

CREATE TABLE character_hq(
    fk_character UUID,
    fk_hq UUID,
    CONSTRAINT pk_character_hq PRIMARY KEY (fk_character,fk_hq),
    FOREIGN KEY (fk_character) REFERENCES character(id),
    FOREIGN KEY (fk_hq) REFERENCES hq(id)
);

CREATE TABLE character_ability(
    fk_character UUID,
    fk_ability UUID,
    CONSTRAINT pk_character_ability PRIMARY KEY (fk_character,fk_ability),
    FOREIGN KEY (fk_character) REFERENCES character(id),
    FOREIGN KEY (fk_ability) REFERENCES ability(id)
);

CREATE TABLE character_game(
    fk_character UUID,
    fk_game UUID,
    CONSTRAINT pk_character_game PRIMARY KEY (fk_character,fk_game),
    FOREIGN KEY (fk_character) REFERENCES character(id),
    FOREIGN KEY (fk_game) REFERENCES game(id)
);

CREATE TABLE team_game(
    fk_team UUID,
    fk_game UUID,
    CONSTRAINT pk_team_game PRIMARY KEY (fk_team,fk_game),
    FOREIGN KEY (fk_team) REFERENCES team(id),
    FOREIGN KEY (fk_game) REFERENCES game(id)
);

CREATE TABLE console_game(
    fk_game UUID,
    fk_console VARCHAR(20)
);

CREATE TABLE character_movie(
    fk_character UUID,
    fk_movie UUID,
    CONSTRAINT pk_character_movie PRIMARY KEY (fk_character,fk_movie),
    FOREIGN KEY (fk_character) REFERENCES character(id),
    FOREIGN KEY (fk_movie) REFERENCES movie(id)
);

