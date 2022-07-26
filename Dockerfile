FROM golang:1.18-alpine

WORKDIR /server

COPY server/go.mod ./
COPY server/go.sum ./

RUN go mod download

COPY server/ ./

RUN go build -o api-server

CMD [ "./api-server" ]