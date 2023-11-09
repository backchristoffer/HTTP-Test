FROM golang:nanoserver-1809
WORKDIR /app
COPY . .
RUN go build
USER default
CMD ["./http-test"]