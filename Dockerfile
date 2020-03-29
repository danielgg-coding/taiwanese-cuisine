FROM golang:1.13

WORKDIR /go/src/github.com/danielgg-coding/taiwanese-cuisine

COPY ./go.mod .
 
RUN go mod download

COPY . .

RUN go build main.go

EXPOSE 8081

# CMD ["go", "run", "main.go"]
CMD ["./main"]