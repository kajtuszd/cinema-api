FROM golang:1.20
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod tidy
RUN go install github.com/codegangsta/gin@latest
RUN go build -o main ./app/main
EXPOSE 8080
CMD ["gin", "--path=/app/app/main", "run", "main.go"]
