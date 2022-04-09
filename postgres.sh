docker run -d \
    --name postgres \
    -e POSTGRES_PASSWORD=password \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v /home/guilherme/projetos/voting/db_data:/var/lib/postgresql/data \
    postgres