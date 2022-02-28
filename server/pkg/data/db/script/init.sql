CREATE TABLE "wifi"
(
  id_users int,
  id_router int,
  id_path int
);

CREATE TABLE "users"
(
    id serial not null unique,
    username CHAR(50) NOT NULL,
    password CHAR(50) NOT NULL
);

CREATE TABLE "router"
(
  id_router serial not null unique,
  idCoordinates int not null,
  transmitterPower int,
  transmitAntenna int,
  receivingAntenna int,
  receiverSensitivity int
);

CREATE TABLE "coordinates_router"
(
  id serial not null unique,
  x INT NOT NULL,
  y INT NOT NULL
);

CREATE TABLE "filePath"
(
    id serial not null unique,
    path_picture char(1000)
);