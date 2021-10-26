FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY .env ./
RUN mkdir ./memestore

RUN go build -o ./memesapi

EXPOSE 8080

CMD [ "./memesapi" ]