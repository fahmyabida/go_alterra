FROM golang:alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go build -o crud-static .

EXPOSE 9000

CMD ["./crud-static"]