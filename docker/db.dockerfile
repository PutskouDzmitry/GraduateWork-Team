FROM postgres:latest

COPY ./server/pkg/data/db/script/init.sql /docker-entrypoint-initdb.d/init.sql

EXPOSE 5432