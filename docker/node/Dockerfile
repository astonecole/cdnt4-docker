FROM node

LABEL MAINTAINER="Vincent Letourneau<vincent@nanoninja.com>"

VOLUME ["/usr/src/app"]
COPY . /usr/src/app
EXPOSE 3000
ENTRYPOINT [ "node", "/usr/src/app/index.js" ]
