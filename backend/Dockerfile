FROM golang:1.26-alpine AS base

WORKDIR /app

COPY go.mod ./
RUN go mod download

FROM base AS dev

RUN go install github.com/air-verse/air@latest

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

FROM base AS builder

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/backend .

FROM alpine:3.24 AS runner

RUN adduser -D -g '' appuser
USER appuser

EXPOSE 8080

COPY --from=builder /bin/backend /bin/backend

ENTRYPOINT ["/bin/backend"]
