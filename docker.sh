# PostgreSQL
docker run -d \
    --name postgres \
    -e POSTGRES_PASSWORD=password \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -p 5432:5432
    -v /home/guilherme/db_data:/var/lib/postgresql/data \
    postgres

# MongoDB
docker run -d \
    --name mongo \
    -e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
    -e MONGO_INITDB_ROOT_PASSWORD=secret \
    -p 27017:27017 \
    -v /home/guilherme/db_data:/data/db \
    --rm \
    mongo

# RabbitMQ
docker run -d \
    --name rabbitmq \
    -p 15672:15672 \
    -p 5672:5672 \
    --rm \
    rabbitmq:3-management