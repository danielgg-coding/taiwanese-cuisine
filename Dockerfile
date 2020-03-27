FROM golang:1.13

WORKDIR /go/src/github.com/danielgg-coding/taiwanese-cuisine

ADD ./go.mod .
 
RUN go mod download

ADD . .

EXPOSE 8080