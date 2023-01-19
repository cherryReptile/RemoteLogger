FROM golang:1.19-alpine AS build

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN apk add git --no-cache

RUN mkdir /tmp/build
RUN mkdir /app
RUN git clone -b master --single-branch https://github.com/cherryReptile/WS-AUTH.git /tmp/build

WORKDIR /tmp/build
RUN go build -o /app/main ./cmd
RUN cp -R migrations /app/migrations
RUN cp ./bash/entrypoint.sh /app/docker-entrypoint.sh
RUN rm -rf /tmp/build
WORKDIR /app

RUN ["chown", "root:root", "/app/docker-entrypoint.sh"]
RUN ["chmod", "+x", "/app/docker-entrypoint.sh"]

ENTRYPOINT ["/app/docker-entrypoint.sh"]

CMD ["./main"]