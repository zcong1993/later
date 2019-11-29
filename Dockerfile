
FROM golang:1.13 AS build
WORKDIR /mnt
ADD go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o ./bin/later main.go

FROM alpine:3.7
WORKDIR /opt
EXPOSE 8080
RUN apk add --no-cache ca-certificates
COPY --from=build /mnt/bin/* /usr/bin/
CMD ["later"]
