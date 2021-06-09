FROM golang:1.16.4
WORKDIR /gocode/src/API
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8000
CMD ["API"]