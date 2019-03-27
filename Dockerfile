FROM golang
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -mod=vendor -o main . 
EXPOSE 80
CMD ["/app/main"]
