# Build consumer
FROM golang:1.21-alpine AS build 

ADD . /app
WORKDIR /app/cmd/consumer
RUN go build -o consumer

# Run in alpine
FROM alpine

COPY --from=build /app/cmd/consumer/ /app/
WORKDIR /app
EXPOSE 8080
ENTRYPOINT [ "./consumer" ]
