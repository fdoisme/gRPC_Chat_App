FROM golang:alpine3.19 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:3.19

COPY --from=build app/main .
EXPOSE 51152
CMD ["./main"]