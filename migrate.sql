CREATE DATABASE myschema;

CREATE TABLE cityInfo
(
    Id SERIAL PRIMARY KEY,
    Name CHARACTER VARYING(30),
    Lat real,
    Lon real,
    Country CHARACTER VARYING(4),
    State CHARACTER VARYING(30)
);

INSERT INTO cityinfo (Name) VALUES 
('Moscow'),
('Saint Petersburg'),
('Novosibirsk'),
('Yekaterinburg'),
('Kazan'),
('Nizhny Novgorod'),
('Chelyabinsk'),
('Samara'),
('Omsk'),
('Rostov-on-Don'),
('Ufa'),
('Krasnoyarsk'),
('Voronezh'),
('Volgograd'),
('Perm'),
('Krasnodar'),
('Saratov'),
('Tyumen'),
('Tolyatti'),
('Izhevsk');

CREATE TABLE cityTemp
(
    city_id int,
    Temp real,
    Date timestamp,
	Full_info jsonb,
	PRIMARY KEY (Temp, Date)
);