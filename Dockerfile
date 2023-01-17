FROM golang:1.19-alpine

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN apk add git --no-cache

RUN mkdir /tmp/build
RUN mkdir /app
RUN git clone -b master --single-branch https://github.com/cherryReptile/WS-AUTH.git /tmp/build

WORKDIR /tmp/build
RUN go build -o /app/main ./cmd
RUN cp -R migrations /app/migrations
RUN cp ./bash/entrypoint.sh /docker-entrypoint.sh
RUN rm -rf /tmp/build
WORKDIR /app

ENTRYPOINT ["/docker-entrypoint.sh"]
RUN ["chown", "root:root", "/docker-entrypoint.sh"]
RUN ["chmod", "+x", "/docker-entrypoint.sh"]

CMD ["./main"]