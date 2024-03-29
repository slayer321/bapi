FROM golang:alpine3.19
COPY . . 
RUN go build -o bapi main.go
EXPOSE 7777
CMD ["./bapi"]