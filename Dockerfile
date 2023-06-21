FROM golang:1.20

WORKDIR /load_database_go
ADD go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./bin/load_database_go .
VOLUME [ "/load_database_go/shared" ]

ENTRYPOINT ["./bin/load_database_go"]
CMD ["/load_database_go/shared/base_teste.txt"]