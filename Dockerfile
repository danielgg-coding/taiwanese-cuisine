# Build the Go API
FROM golang:1.13 AS go_builder
ADD . /app
WORKDIR /app/backend
RUN go mod download

# RUN go build main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

# Build the React application
FROM node:alpine AS node_builder
COPY --from=go_builder /app/frontend ./
RUN npm install
RUN npm run build

# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=go_builder /main ./
COPY --from=node_builder /build ./web
RUN chmod +x ./main

EXPOSE 8081
CMD ["./main"]

# COPY . .