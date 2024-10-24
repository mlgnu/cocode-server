FROM golang:latest
#AS builder

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod verify
# RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .
# RUN goose up

# RUN CGO_ENABLED=0 go build -o cocode-server

# FROM alpine:latest
# WORKDIR /app
RUN go install github.com/air-verse/air@latest

# COPY --from=builder /app/cocode-server /app/cocode-server


EXPOSE 8080
CMD ["air", "-c", ".air.toml"]
# CMD ["./cocode-server"]
