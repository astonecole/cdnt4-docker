# Step one
FROM golang:alpine as app

LABEL MAINTAINER="Vincent Letourneau <vincent@nanoninja.com>"
WORKDIR /go/src/app
RUN apk add --update git \
    && go get github.com/nanoninja/bulma
COPY server.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo -o main server.go

# Step two
FROM scratch
COPY --from=app /go/src/app .
EXPOSE 3000
CMD [ "./main" ]
