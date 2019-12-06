FROM golang:alpine
RUN apk --no-cache add \
    build-base \
    git \
    bzr \
    mercurial \
    gcc

ADD . /src
RUN cd /src && \
    GOOS=linux go \
    build -o lne-user \
    cmd/main.go

ENTRYPOINT ./lne-user start