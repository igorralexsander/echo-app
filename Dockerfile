FROM golang:1.22 as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o echo-app cmd/echo-app/main.go

# Execution container
FROM alpine
WORKDIR /app
EXPOSE 8080
COPY --from=builder /build/echo-app /app/echo-app
ENTRYPOINT ["/app/echo-app"]

