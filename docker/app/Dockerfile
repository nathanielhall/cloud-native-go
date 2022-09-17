# Build environment
# -----------------
FROM golang:1.18-alpine as build-env
WORKDIR /cloud-native-go
RUN apk update && apk add --no-cache gcc musl-dev git bash
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash
COPY --from=build-env /cloud-native-go/bin/app /cloud-native-go/

EXPOSE 8080
CMD ["/cloud-native-go/app"]