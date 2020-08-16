FROM golang:1.14

RUN mkdir app/
WORKDIR /app/

ADD . ./
RUN go build main.go

COPY . . 
CMD ["./main"]