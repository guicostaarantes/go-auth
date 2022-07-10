FROM golang:1.18.0-alpine3.14 AS build
WORKDIR /app
RUN apk add --no-cache build-base
RUN go install github.com/go-delve/delve/cmd/dlv@v1.8.3
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go install github.com/99designs/gqlgen
COPY graph/ graph
COPY modules/ modules
COPY utils/ utils
COPY scripts/ scripts
COPY gqlgen.yml .
COPY main.go .
RUN chmod -R +x scripts
RUN scripts/generate.sh
RUN go build -gcflags="all=-N -l" -o /out/main main.go
CMD ["/go/bin/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "--continue", "/out/main"]
