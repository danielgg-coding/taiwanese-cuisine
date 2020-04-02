FROM golang:1.13

WORKDIR /go/src/github.com/danielgg-coding/taiwanese-cuisine/

COPY ./backend/go.mod .
 
RUN go mod download

COPY . .

RUN go build backend/main.go

EXPOSE 8081

CMD ["./main"]