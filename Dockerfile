FROM golang:1.18-alpine
WORKDIR /app
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64
COPY ./ ./
RUN go mod download
RUN go build cmd/app/parser.go
CMD [ "./parser" ]