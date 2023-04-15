FROM golang:1.20-alpine
WORKDIR /cloud-native-go

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/api \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

CMD ["/cloud-native-go/bin/api"]
EXPOSE 8080