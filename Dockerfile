FROM golang:latest
RUN mkdir /app
workdir /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV PORT 8080
RUN go build
CMD ["./app/main.exe"]