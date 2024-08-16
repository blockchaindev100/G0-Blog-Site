FROM golang:1.22.6-alpine AS builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /blogsite
EXPOSE 8080
CMD [ "/blogsite" ]