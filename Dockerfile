FROM golang:1.16

RUN apk add --no-cache git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./main/main.go .

EXPOSE 80

CMD [ "/main/main.go" ]
