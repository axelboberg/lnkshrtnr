FROM golang:1.15

WORKDIR /go/src/github.com/axelboberg/lnkshrtnr
COPY . .

# Install dependencies
# and compile the app
RUN go get -d -v ./...
RUN go install
RUN go build main.go

CMD ["./main"]