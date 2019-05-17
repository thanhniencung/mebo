FROM golang:alpine

COPY . .

CMD ["./cmd/cmd"]

EXPOSE 3000