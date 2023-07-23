FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

RUN echo $PWD && ls

RUN go mod download 
RUN go mod verify

RUN go build -v -o binary

ENTRYPOINT ["/app/binary"]