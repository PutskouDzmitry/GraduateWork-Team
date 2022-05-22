CREATE TABLE "wifi_data_models"
(
  id SERIAL NOT NULL UNIQUE,
  id_user_data int,
  id_router_wifi int,
  id_file_path int
);

CREATE TABLE "file_path"
(
    id SERIAL NOT NULL UNIQUE,
    path_input char(90),
    path_output char(90)
);

CREATE TABLE "users"
(
    id SERIAL NOT NULL UNIQUE,
    username CHAR(50) NOT NULL,
    password CHAR(100) NOT NULL
);

CREATE TABLE "router_data_models"
(
  id_router SERIAL NOT NULL UNIQUE,
  coordinates_of_router_id int not null,
  transmitter_power float,
  gain_of_transmitting_antenna float,
  gain_of_receiving_antenna    float,
  speed                     int,
  signal_loss_transmitting    float,
  signal_loss_receiving float,
  number_of_channels float,
  typeOfSignal float,
  scale float,
  com float
);

CREATE TABLE "coordinates_points"
(
  id SERIAL NOT NULL UNIQUE,
  x INT NOT NULL,
  y INT NOT NULL
);