FROM golang:1.23.6-alpine3.20
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]