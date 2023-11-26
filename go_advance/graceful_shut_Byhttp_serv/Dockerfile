# syntax=docker/dockerfile:1

FROM golang:1.21.1

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /graceful_shut_Byhttp_serv

EXPOSE 8080

# Run

CMD ["/graceful_shut_Byhttp_serv"]