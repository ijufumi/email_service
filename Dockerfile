FROM golang:1.19-bullseye as build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main /app/cmd/app/

FROM debian:bullseye-slim as deploy

WORKDIR /app

COPY --from=build /app/main /app/main

EXPOSE 8080

ENTRYPOINT ["/app/main"]