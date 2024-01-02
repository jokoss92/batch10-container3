FROM golang:alpine as build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go .
RUN go build -o /go-app

FROM alpine:latest
WORKDIR /
COPY --from=build /go-app /go-app
EXPOSE 8000
ENTRYPOINT ["/go-app"]

