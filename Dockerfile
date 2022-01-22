FROM golang:1.16-alpine

WORKDIR /app

COPY FindTheWord ./

RUN go build /FindTheWord/main.go

RUN ./FindTheWord/main input.txt

CMD [ "cat", "response.txt" ]