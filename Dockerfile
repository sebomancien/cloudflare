FROM golang:1.23-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app /src/...

# Run the tests in the container
FROM builder AS tester
RUN go test -v /src/...

# Deploy the application binary into a lean image
FROM alpine:latest AS releaser

WORKDIR /app

COPY --from=builder /app /bin

RUN addgroup -S nonroot && adduser -S nonroot -G nonroot

USER nonroot:nonroot

ENTRYPOINT ["cloudflare", "records", "update"]
