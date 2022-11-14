## Getting started with Postgres

1) Download [Docker Desktop](https://docs.docker.com/engine/install/)
2) Open terminal in folder with `compose.yml` and run `docker-compose --env-file ./config/.env.dev up -d`.
3) After container start you can directly access the db through the pgadmin panel `http://localhost:5050/`
    - In pgadmin you need to set `general` settings (`name`) and `connection` (`hostname`,`db_name`, `db_user`, `db_password`)
4) Or you can connect throw IDE
    - IDE will ask for the localhost of the database (`http://localhost:5432/`) and passwords for it (`db_name`, `db_user`, `db_password`).

`docker-compose stop` - stop containers

`docker-compose down` - stop and remove containers