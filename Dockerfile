FROM golang:1.16

WORKDIR /main

COPY . .

EXPOSE 8080

CMD ["go", "run", "main/main.go"]
