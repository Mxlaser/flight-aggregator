FROM golang:1.22-alpine

ENV GOTOOLCHAIN=auto
# Toolchain C pour CGO et -race
RUN apk add --no-cache make git curl build-base

# Activer CGO pour les tests avec -race
ENV CGO_ENABLED=1

RUN go install github.com/air-verse/air@latest
RUN go install gotest.tools/gotestsum@latest

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]
