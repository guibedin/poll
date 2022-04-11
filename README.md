# voting
Inspired by the [Voting System Challenge](https://dev.to/zanfranceschi/desafio-sistema-de-votacao-bbb-50e3) by [Zan Franceschi](https://dev.to/zanfranceschi).

# Create your poll


# Local Setup
### Run PostgreSQL
```
docker run -d \
    --name postgres \
    -e POSTGRES_PASSWORD=password \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v ${DB_DATA_PATH}:/var/lib/postgresql/data \
    postgres
```
### Run RabbitMQ
```
docker run -d \
    --name rabbitmq \
    -p 15672:15672 \
    -p 5672:5672 \
    rabbitmq:3-management
```
http://localhost:15672/ - guest/guest
