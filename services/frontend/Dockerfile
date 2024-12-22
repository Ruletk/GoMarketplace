FROM node:18-alpine as build

WORKDIR /app

COPY frontend/package.json ./

RUN npm install

COPY frontend/ .

RUN npm run build


FROM nginx:1.27.3

COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]