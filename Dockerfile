FROM golang:1.19-alpine

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN apk add git --no-cache

RUN mkdir /tmp/build
RUN mkdir /app
RUN git clone -b v1.1.0 --single-branch https://github.com/cherryReptile/WS-AUTH.git /tmp/build

WORKDIR /tmp/build
RUN go build -o /app/main ./cmd
RUN rm -rf /tmp/build
WORKDIR /app
RUN ls -la