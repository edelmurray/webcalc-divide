FROM golang:1.16-alpine

WORKDIR /main

COPY . .

EXPOSE 8080

CMD ["go", "run", "main/main.go"]
