FROM golang:1.13 as builder
RUN mkdir /code
ADD . /code
WORKDIR /code
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

########## Start a new stage ##########
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
RUN mkdir /root/assets
COPY --from=builder /code/wait-for.sh /root/
COPY --from=builder /code/main .
COPY --from=builder /code/assets/ /root/assets/
CMD ["./main"]