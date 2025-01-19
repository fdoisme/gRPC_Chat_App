FROM golang:alpine3.19 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o auth-service ./auth-service/.

FROM alpine:3.19
WORKDIR /auth-service/
COPY --from=build app/auth-service .
EXPOSE 51151
CMD ["./auth-service"]