docker run -d \
    --name postgres \
    -e POSTGRES_PASSWORD=password \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v /home/guilherme/projetos/voting/db_data:/var/lib/postgresql/data \
    postgres

docker run -d \
    --name rabbitmq \
    -p 15672:15672 \
    -p 5672:5672 \
    rabbitmq:3-management