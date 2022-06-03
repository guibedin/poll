# Poll
Inspired by the [Voting System Challenge](https://dev.to/zanfranceschi/desafio-sistema-de-votacao-bbb-50e3) by [Zan Franceschi](https://github.com/zanfranceschi).

### Resources
- Mat Ryer - [How I build HTTP services after eight years](https://github.com/matryer/2019-talks/tree/master/Mat%20Ryer%20-%20How%20I%20build%20HTTP%20services%20after%20eight%20years)
- Kat Zien - [Go structure examples](https://github.com/katzien/go-structure-examples)

# Next steps
- **Tests** - Write tests for this project.
- **Front End** - A front end using this application.
- **Repositories** - Make a MongoDB and a File repository implementation.

# Create your poll

# How it works
This is Poll creation system split up in two parts. One is the server, which receives all requests through HTTP.

When the `server` receives a request to vote, it sends a message to a RabbitMQ queue and returns `HTTP 202 - ACCEPTED` to the client.

The `consumer` is responsible for reading messages from RabbitM, and then updating the Database counting the vote.

## Routes
###  **Get All Polls** - `GET /api/polls`

Example request:
```sh
curl --location --request GET 'localhost:8080/api/polls/'
```
Response body:
```json
[
    {
        "id": 1,
        "title": "Poll 1",
        "is_active": true,
        "is_multiple_choice": true,
        "options": [
            {
                "id": 1,
                "title": "Option 1",
                "votes": 9,
                "created_at": "2022-05-26T23:26:31.009656Z",
                "updated_at": "2022-05-26T23:26:31.009656Z"
            },
            {
                "id": 2,
                "title": "Option 2",
                "votes": 5,
                "created_at": "2022-05-26T23:26:31.013177Z",
                "updated_at": "2022-05-26T23:26:31.013177Z"
            }
        ],
        "created_at": "2022-05-26T23:26:31.0062Z",
        "updated_at": "2022-05-26T23:26:31.0062Z"
    },
    {
        "id": 2,
        "title": "Poll 2",
        "is_active": false,
        "is_multiple_choice": false,
        "options": [
            {
                "id": 3,
                "title": "Option 3",
                "votes": 5,
                "created_at": "2022-05-26T23:26:41.997339Z",
                "updated_at": "2022-05-26T23:26:41.997339Z"
            },
            {
                "id": 4,
                "title": "Option 4",
                "votes": 5,
                "created_at": "2022-05-26T23:26:42.001372Z",
                "updated_at": "2022-05-26T23:26:42.001372Z"
            }
        ],
        "created_at": "2022-05-26T23:26:41.993744Z",
        "updated_at": "2022-05-26T23:26:41.993744Z"
    }
]
```

### **Create New Poll** - `POST /api/polls`

Example Request:
```sh
curl --location --request POST 'localhost:8080/api/polls' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Poll 1",
    "is_active": true,
    "is_multiple_choice": true,
    "options": [
        {
            "title": "Option 1"
        },
        {
            "title": "Option 2"
        }
    ]
}'
```
Request body:
```json
{
    "title": "Poll 1",
    "options": [
        {
            "title": "Option 1"
        },
        {
            "title": "Option 2"
        }
    ]
}
```
Response body:
```json
{
    "poll_id": 1
}
```

### **Get Poll By ID** - `GET /api/polls/:id`

Example request:
```sh
curl --location --request GET 'localhost:8080/api/polls/1'
```

Response body:
```json
{
    "id": 1,
    "title": "Poll 1",
    "is_active": true,
    "is_multiple_choice": true,
    "options": [
        {
            "id": 1,
            "title": "Option 1",
            "votes": 9,
            "created_at": "2022-05-26T23:26:31.009656Z",
            "updated_at": "2022-05-26T23:26:31.009656Z"
        },
        {
            "id": 2,
            "title": "Option 2",
            "votes": 5,
            "created_at": "2022-05-26T23:26:31.013177Z",
            "updated_at": "2022-05-26T23:26:31.013177Z"
        }
    ],
    "created_at": "2022-05-26T23:26:31.0062Z",
    "updated_at": "2022-05-26T23:26:31.0062Z"
}
```

### **Single Option Vote** - `POST /api/polls/:id/vote`

Example request:
```sh
curl --location --request POST 'localhost:8080/api/polls/1/vote' \
--header 'Content-Type: application/json' \
--data-raw '{    
    "voter": "voter",
    "option_id": 1
}'
```

### **Multiple Options Vote** - `POST /api/polls/:id/votes`

Example request:
```sh
curl --location --request POST 'localhost:8080/api/polls/2/votes' \
--header 'Content-Type: application/json' \
--data-raw '{    
    "voter": "voter",
    "option_ids": [3, 4]
}'
```

# Local Setup

## PostgreSQL

### Create psql docker image

From inside the `db` folder, run: `docker build . -t <tag>`

Then, run the image setting the variables as you wish:
```
docker run -d \
    --name postgres \
    -e POSTGRES_USER=<user> \
    -e POSTGRES_PASSWORD=<password> \
    -e POSTGRES_DB=poll \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -p 5432:5432 \
    -v <local path for your db data>:/var/lib/postgresql/data \
    --rm \
    <tag>
```

## RabbitMQ
```
docker run -d \
    --name rabbitmq \
    -p 15672:15672 \
    -p 5672:5672 \
    rabbitmq:3-management
```
To see your RabbitMQ dashboard, access http://localhost:15672/ with the credentials `guest/guest`

## Server
`go run cmd/web/main.go`

## Consumer
`go run cmd/consumer/main.go`