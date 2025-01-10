FROM golang:latest as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/main.go


FROM gcr.io/distroless/static as runner

WORKDIR /app

COPY --from=builder /app/main /app

ENTRYPOINT ["/app/main"]