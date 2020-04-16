FROM node:alpine AS build_stage
WORKDIR /app
COPY ./package.json .
RUN npm install
COPY . .
RUN npm run build

FROM nginx
EXPOSE 3000
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
COPY --from=build_stage /app/build /usr/share/nginx/html/taiwanese-cuisine-frontend