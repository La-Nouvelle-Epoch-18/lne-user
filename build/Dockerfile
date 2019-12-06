# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
RUN cd /src && GOOS=linux GOARCH=amd64 go build -o lne-user cmd/main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/lne-user /app/
ENTRYPOINT ./app/lne-user