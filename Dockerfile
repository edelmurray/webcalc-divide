FROM golang:1.16-alpine

RUN apk add --no-cache git

COPY go.mod .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./main/main.go .

EXPOSE 80

CMD [ "/main/main.go" ]
