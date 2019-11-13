FROM golang:1.12 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /usr/src/app
COPY . .
RUN go build

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/src/app/devops-practice /devops-practice
EXPOSE 8080
CMD ["/devops-practice"]
