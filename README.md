# voting
Inspired by the [Voting System Challenge](https://dev.to/zanfranceschi/desafio-sistema-de-votacao-bbb-50e3) by [Zan Franceschi](https://dev.to/zanfranceschi).

# Create your poll


# Local Setup
### Run MongoDB
```
docker run -d \
    --name mongo \
    -e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
    -e MONGO_INITDB_ROOT_PASSWORD=secret \
    -p 27017:27017 \
    -v /home/guilherme/projetos/voting/db_data:/data/db \
    mongo
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
