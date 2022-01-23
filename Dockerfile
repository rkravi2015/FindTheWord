FROM golang:1.16-alpine

WORKDIR /app

COPY * ./

RUN go build main.go

RUN ./main input.txt

CMD [ "cat", "response.txt" ]