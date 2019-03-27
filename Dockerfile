FROM golang
RUN mkdir /app 
ADD . /app/ 
EXPOSE 80
CMD ["/app/main"]
