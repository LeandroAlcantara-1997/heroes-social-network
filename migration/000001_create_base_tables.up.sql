CREATE TABLE character(
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    character_name VARCHAR(100),
    civil_name VARCHAR(100),
    heroe BOOLEAN
);


CREATE TABLE super_power(
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    description VARCHAR(300)
);

CREATE TABLE character_super_power(
    fk_character UUID,
    fk_super_power UUID,
    CONSTRAINT pk_character_super_power PRIMARY KEY (fk_character,fk_super_power),
    FOREIGN KEY (fk_character) REFERENCES character(id),
    FOREIGN KEY (fk_super_power) REFERENCES super_power(id)
);
