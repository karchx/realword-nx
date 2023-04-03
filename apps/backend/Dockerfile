# STEP 1: Modules caching
FROM golang:1.20-alpine AS modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# STEP 2: Builder
FROM golang:1.20-alpine AS builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app .

# STEP 3: Final
FROM scratch

EXPOSE 5001

# GOPATH for scratch image is /
COPY --from=builder /bin/app /app
CMD [ "/app" ]