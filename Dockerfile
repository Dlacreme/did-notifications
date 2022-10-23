FROM golang:alpine
RUN mkdir /app
RUN go install github.com/joho/godotenv/cmd/godotenv@latest
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
