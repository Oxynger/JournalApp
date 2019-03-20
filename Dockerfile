FROM golang:latest
# Возможно понадобиться добавить в PATH go/bin
ADD . /code
WORKDIR /code
RUN go get
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init
RUN go build -o app ./main.go
CMD app