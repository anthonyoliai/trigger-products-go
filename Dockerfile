FROM golang:1.23-alpine
RUN apk add build-base

WORKDIR /app

ADD . /app
RUN go build -o /products-trigger

CMD [ "/products-trigger" ]
