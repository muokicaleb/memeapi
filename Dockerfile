FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN mkdir /memestore

RUN go build -o /memes

EXPOSE 8080

CMD [ "/memes" ]