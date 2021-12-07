FROM golang:1.17-buster

ENV GOPROXY="https://goproxy.cn,direct"

ENV GOPRIVATE="devgit.westwell-lab.info"

COPY .netrc /root/.netrc

RUN go get -u github.com/swaggo/swag/cmd/swag
