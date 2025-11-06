FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG BUILD_VERSION="dev"
ARG BUILD_TIME="unknown"
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -trimpath -ldflags "-s -w -X main.version=${BUILD_VERSION} -X main.buildTime=${BUILD_TIME}" -o /app/bin/todo-backend .

FROM scratch AS runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /app/bin/todo-backend /usr/local/bin/todo-backend

USER 1000

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/todo-backend"]
