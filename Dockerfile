FROM golang:alpine
WORKDIR /app
COPY . .
RUN go get http-test
RUN go build
USER default
CMD ["./http-test"]
