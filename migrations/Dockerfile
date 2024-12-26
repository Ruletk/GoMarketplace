FROM golang:latest as builder

WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.24.0

COPY . .

CMD ["goose", "up", "&&", "echo", "Migrations completed! I am done! 🎉"]

