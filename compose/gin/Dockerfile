FROM golang:1.14.4-alpine3.12

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/src/drizzle
COPY . .
COPY compose/gin/app.ini  conf/

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["drizzle"]