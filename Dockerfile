FROM golang:1.19-alpine

RUN apk update && apk add curl curl-dev gcc musl-dev

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o . ./...

EXPOSE 8000

CMD [ "./currency" ]
