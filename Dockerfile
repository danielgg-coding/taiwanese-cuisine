FROM node:alpine AS build_front_stage
WORKDIR /app
COPY ./frontend/package.json .
RUN npm install
COPY ./frontend .
RUN npm run build

# FROM nginx
# EXPOSE 3000
# COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
# COPY --from=build_front_stage /app/build /usr/share/nginx/html

# syntax=docker/dockerfile:1.0.0-experimental
FROM golang:1.13 AS build_back_stage
WORKDIR /app
COPY ./backend/go.mod .
RUN go mod download
COPY ./backend .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o main

FROM alpine
COPY --from=build_back_stage /app/main ./
COPY --from=build_back_stage /app/serviceAccountKey ./
COPY --from=build_front_stage /app/build  ./build

EXPOSE 8080

CMD ["/main"]