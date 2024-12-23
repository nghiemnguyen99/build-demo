FROM golang:1.23.0-alpine

#for migration
#`mariadb-connector-c` is needed to suppress the error `ERROR 1045 (28000): Plugin caching_sha2_password could not be loaded: Error loading shared library /usr/lib/mariadb/plugin/caching_sha2_password.so: No such file or directory`.
# RUN apk add mysql-client mariadb-connector-c

# RUN adduser -D user
# USER user
ENV GOPATH=/home/user/go

RUN go install github.com/air-verse/air@latest
ENV PATH ${PATH}:/home/user/go/bin

#downloads the dependencies
WORKDIR /go_basic/
COPY ./ ./
RUN go mod download

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# USER root
WORKDIR /go_basic/cmd/
RUN go build -o main
RUN chmod +x ./main

# WORKDIR /go_basic/
# COPY cmd/build ./

CMD ["air"]

