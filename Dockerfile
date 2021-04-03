FROM golang:1.16-buster

WORKDIR /go/src/github.com/task4233/slide-decks

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 3000
CMD ["present", "-http", "0.0.0.0:3000", "-orighost", "localhost"]
