FROM golang:1.16-alpine

<<<<<<< HEAD
RUN apk add --no-cache git

COPY go.mod .

RUN go mod download
=======
WORKDIR /main
>>>>>>> 8773d6d05ccfdfe5999cf32ac34a4a16a75a321c

COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...
#RUN chmod -R 755 ./main
EXPOSE 8080

CMD ["go", "run", "main/main.go"]
