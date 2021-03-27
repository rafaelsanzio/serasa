FROM golang:1.15

WORKDIR /
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build main.go

CMD ./main
