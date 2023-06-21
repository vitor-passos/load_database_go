FROM golang:1.20

WORKDIR /load_database_go
ADD go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./bin/load_database_go .

ENTRYPOINT ["./bin/load_database_go"]
CMD ["base_teste.txt"]