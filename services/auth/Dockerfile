FROM golang:latest as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build main.go


FROM gcr.io/distroless/static as runner

COPY --from=builder /app/main /

ENTRYPOINT ["/main"]