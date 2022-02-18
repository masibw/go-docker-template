FROM golang:1.17.6

ENV CGO_ENABLED=0

WORKDIR /workdir

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build .

CMD ["/workdir/go-docker-template"]