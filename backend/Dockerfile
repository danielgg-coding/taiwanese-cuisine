# syntax=docker/dockerfile:1.0.0-experimental
FROM golang:1.13 AS build_stage
WORKDIR /app
COPY ./go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o main

FROM alpine
COPY --from=build_stage /app/main ./
COPY --from=build_stage /app/serviceAccountKey ./
CMD ["/main"]