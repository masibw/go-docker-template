FROM golang:1.17.6 AS build

ENV CGO_ENABLED=0

WORKDIR /workdir

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app .

FROM gcr.io/distroless/static-debian11:latest

WORKDIR /workdir

COPY --from=build /workdir/app /workdir/

CMD ["/workdir/app"]