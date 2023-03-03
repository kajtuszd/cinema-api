FROM golang:1.20
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod tidy
RUN go install github.com/codegangsta/gin@latest
RUN go build -o main ./app
EXPOSE 8080
CMD ["gin", "--path=/app/app", "run", "main.go"]
