FROM golang:1.18.0-alpine3.14 AS build
WORKDIR /app
RUN apk add --no-cache build-base
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY graph/ graph
COPY modules/ modules
COPY utils/ utils
COPY scripts/ scripts
COPY gqlgen.yml .
COPY main.go .
COPY tools.go .
RUN chmod -R +x scripts
RUN scripts/generate.sh
RUN scripts/build.sh

FROM alpine:3.14
COPY --from=build /out /
CMD ["/main"]
