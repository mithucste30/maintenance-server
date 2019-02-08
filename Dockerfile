FROM golang:alpine as server_build

# Add build deps
RUN apk add --update gcc g++ git

COPY . /appbuild

RUN set -ex \
    && go version \
    && cd /appbuild \
    && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o server \
    && rm -rf *.mod *.go

# Build deployable server
FROM alpine:latest

WORKDIR /app

COPY --from=server_build /appbuild /app/

EXPOSE 80

CMD ["./server"]
