FROM golang:latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["bash"]