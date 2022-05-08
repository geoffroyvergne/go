FROM golang

COPY . /go/src/gin-starter-redis
WORKDIR /go/src/gin-starter-redis

RUN go get -d ./.
RUN go build -o app

EXPOSE 3000
CMD ["./app"]