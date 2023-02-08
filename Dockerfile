FROM golang

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-pg.sh

RUN go mod download
RUN go build -o go-api ./cmd/main.go

CMD ["./go-api"]