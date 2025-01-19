FROM golang:alpine3.19 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o api-gateway ./api-gateway/.

FROM alpine:3.19
WORKDIR /api-gateway/
COPY --from=build app/api-gateway .
EXPOSE 8181
CMD ["./api-gateway"]