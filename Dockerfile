# Use the official Go image as the base
FROM golang:1.22 as base

# Install the Air binary for live code reloading
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Set the working directory inside the container
WORKDIR /app

COPY ./app .

EXPOSE 80/tcp

RUN go get
RUN go build -o bin /app/.

ENTRYPOINT [ "/app/bin" ]
