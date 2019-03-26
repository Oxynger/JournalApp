FROM golang
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get github.com/swaggo/swag/cmd/swag \
    && swag init \
    && go build -o main . 
EXPOSE 80
CMD ["/app/main"]
