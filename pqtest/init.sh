#!/usr/bin/env bash

# TODO: args for docker-image name + port + dbname
# TODO: check if docker is already running (config should be the same) -> if so, clean only database (faster)

echo "Remove old container..."
docker rm -f -v postgres5432

echo "Create new postgres docker container..."
docker run -d --name postgres5432 -p 5432:5432 -e POSTGRES_USER=postgres postgres:latest

echo "Checking Postgres availability..."
while : ; do
    echo "Waiting for Postgres..."
    sleep 3
    upstartCheck=$(docker exec postgres5432 /bin/sh -c "ps aux | grep 'postgres' | grep 'docker-entrypoint.sh' | grep -v 'grep'")
    [[ $upstartCheck == *"docker-entrypoint.sh"* ]] || break
done

echo "Create database..."
docker exec postgres5432 /bin/sh -c "su postgres --command 'createdb -O postgres test_pqx'"