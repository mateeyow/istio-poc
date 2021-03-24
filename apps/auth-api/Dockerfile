FROM golang:1.15-alpine as build

WORKDIR /src

COPY go.sum /src
COPY go.mod /src

COPY main.go /src

RUN go build -o main

FROM alpine:3.12

RUN mkdir -p /app
COPY --from=build /src/main /app/main
CMD [ "./app/main" ]