FROM golang:1.16-alpine

WORKDIR $GOPATH/src/FindTheWord

COPY * ./

RUN go build /main.go

RUN ./main input.txt

CMD [ "cat", "response.txt" ]
