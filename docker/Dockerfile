FROM golang:1.23-alpine

ARG SERVICE

WORKDIR /app/service/${SERVICE}

RUN go install github.com/air-verse/air@latest

COPY ./service/${SERVICE}/go.mod ./service/${SERVICE}/go.sum ./
COPY ./utils/ /app/utils/
RUN go mod download

CMD ["sh", "-c", "air", "-c", ".air.toml"]
