
FROM golang:1.16-alpone as builder

WORKDIR /build

COPY go.mod go.sum ./

COPY . .

EXPOSE 3000

ENTRYPOINT ["/"]
