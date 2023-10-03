# Build application
FROM golang:1.21-alpine AS build

ADD . /app
WORKDIR /app/cmd/web
RUN go build -o web

# Run in alpine container
FROM alpine

COPY --from=build /app/cmd/web/ /app/
WORKDIR /app
ENTRYPOINT [ "./web" ]
