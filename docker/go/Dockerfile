FROM golang:latest

WORKDIR /go/src/github.com/ukkary/repo_name
ADD . /go/src/github.com/ukkary/repo_name

RUN apt-get update \
  && go get -u github.com/gin-gonic/gin \
  && go get -u github.com/go-sql-driver/mysql

EXPOSE 3003