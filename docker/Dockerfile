FROM golang:1.23-alpine

ARG SERVICE

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY ./service/${SERVICE}/go.mod ./service/${SERVICE}/go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
