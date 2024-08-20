FROM golang:1.23

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go ./

RUN go build -buildvcs=false -o ./request-logger
RUN ls -alh

EXPOSE 8080

CMD ["./request-logger"]

