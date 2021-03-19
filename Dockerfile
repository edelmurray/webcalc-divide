FROM golang:1.16

WORKDIR /main

COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...
#RUN chmod -R 755 ./main
EXPOSE 8080

CMD ["go", "run", "main/main.go"]
