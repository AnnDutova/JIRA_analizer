## Getting started with Postgres

1) Download [Docker Desktop](https://docs.docker.com/engine/install/)
2) Open terminal in folder with `compose.yml` and run `docker-compose --env-file ./config/.env.dev up -d`.
3) After container start you can directly access the db through the pgadmin panel `http://localhost:5050/`
    - In pgadmin you need to set `general` settings (`name`) and `connection` (`hostname`,`db_name`, `db_user`, `db_password`)
4) Or you can connect throw IDE
    - IDE will ask for the localhost of the database (`http://localhost:5432/`) and passwords for it (`db_name`, `db_user`, `db_password`).

`docker-compose stop` - stop containers

`docker-compose down` - stop and remove containers

## Getting started with k8s
Start:
```yaml
kubectl apply -f service.yaml
kubectl apply -f config-map-db.yaml
kubectl create configmap master-slave-config --from-file=$(PROJECT_PATH)/db-init-k8s/ --from-file=$(PROJECT_PATH)/db-init-scripts/
kubectl apply -f statefull-set.yaml
```
Delete:
```yaml
 kubectl delete service/db-service
 kubectl delete configmap/db-config
 kubectl delete configmap/master-slave-config
 kubectl delete statefulset.apps/postgres-db
 kubectl delete pvc --all
 kubectl delete pv --all
```
Command in CLI for check:
```yaml
psql -U pguser -l
psql --port=5432 --user=pguser testdb
\dt
su
psql -U postgres
\dt
\df
\x
Insert into author(name) values('some name');
```