# Poll
Inspired by the [Voting System Challenge](https://dev.to/zanfranceschi/desafio-sistema-de-votacao-bbb-50e3) by [Zan Franceschi](https://github.com/zanfranceschi).

# Next steps
- **Tests** - Write tests for this project.
- **Repositories** - Make a MongoDB and a File repository implementation

# Create your poll

# How it works - **The description bellow is a bit outdated, will work on it in the next commits.**
This is Poll creation system split up in two parts. One is the server, which receives all requests through HTTP.

When the `server` receives a request to vote, it sends a message to a RabbitMQ queue and returns `HTTP 202 - ACCEPTED` to the client.

The `worker` is responsible for receiving that message, reading it from RabbitMQ, and then updating the Database counting the vote.

## Routes
###  **Get All Polls**
`GET  /api/polls`

Response body:
```json
[
    {
        "poll": {
            "poll_id": 1,
            "title": "Poll 1",
            "is_active": true,
            "created_on": "2022-04-23T10:43:46.414589Z"
        },
        "options": [
            {
                "option_id": 1,
                "poll_id": 1,
                "title": "Option 1",
                "votes": 0
            },
            {
                "option_id": 2,
                "poll_id": 1,
                "title": "Option 2",
                "votes": 0
            }
        ]
    },
    {
        "poll": {
            "poll_id": 2,
            "title": "Poll 2",
            "is_active": true,
            "created_on": "2022-04-23T10:45:01.722008Z"
        },
        "options": [
            {
                "option_id": 3,
                "poll_id": 2,
                "title": "Option 3",
                "votes": 0
            },
            {
                "option_id": 4,
                "poll_id": 2,
                "title": "Option 4",
                "votes": 0
            },
            {
                "option_id": 5,
                "poll_id": 2,
                "title": "Option 5",
                "votes": 0
            }
        ]
    }
]
```

### **Create New Poll**
`POST /api/polls`

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

### **Get Poll By ID**
`GET  /api/polls/:id`

Response body:
```json
{
    "poll": {
        "poll_id": 1,
        "title": "Poll 1",
        "is_active": true,
        "created_on": "2022-04-23T10:43:46.414589Z"
    },
    "options": [
        {
            "option_id": 1,
            "poll_id": 1,
            "title": "Option 1",
            "votes": 0
        },
        {
            "option_id": 2,
            "poll_id": 1,
            "title": "Option 2",
            "votes": 0
        }
    ]
}
```

### **Vote**
`POST /api/polls/:id/vote`

Request body:
```json
{    
    "option_ids": [1, 2]
}
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
`go run cmd/server/main.go`

## Worker
`go run cmd/worker/main.go`