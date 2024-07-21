FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/skinaapis ./api/main.go

FROM alpine:latest AS final

WORKDIR /app

COPY --from=build /app/bin/skinaapis ./

EXPOSE 9090

ENTRYPOINT [ "./skinaapis" ]