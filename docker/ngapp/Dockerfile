FROM node:12.14.1-alpine as ngapp

LABEL MAINTAINER="Vincent Letourneau<vincent@nanoninja.com>"
WORKDIR /usr/src/app
COPY . .
RUN apk add --update python python3 \
    && npm cache clean --force \
    && npm i --no-optionnal \
    && npm run build

# Etape 2
FROM nginx:alpine
COPY --from=ngapp /usr/src/app/dist/ngapp /usr/share/nginx/html
