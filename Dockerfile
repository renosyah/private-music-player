# docker file for receptionist-core-api app
FROM golang:latest as builder
ADD . /go/src/github.com/renosyah/receptionist-core-api
WORKDIR /go/src/github.com/renosyah/receptionist-core-api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
RUN rm -rf /vendor
CMD ./main --config=.heroku.toml
MAINTAINER syahputrareno975@gmail.com
