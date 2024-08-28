FROM golang:1.22.6-alpine3.20 as builder

WORKDIR /app
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl
RUN adduser -D -g '' golang

WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env
COPY docs/swagger.json docs/swagger.json

USER golang

EXPOSE 8000

CMD ["/app/main"]