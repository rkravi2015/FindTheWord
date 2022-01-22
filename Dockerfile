FROM golang:1.16-alpine

WORKDIR /app

RUN go mod download

COPY * ./

RUN go build -o /search-word

RUN ./search-word input.txt

CMD [ "cat", "response.txt" ]