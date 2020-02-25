FROM golang:1.13 as builder
RUN mkdir /code
ADD . /code
WORKDIR /code
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

########## Start a new stage ##########
FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir -p /code/assets && mkdir -p /code/html
WORKDIR /code
COPY --from=builder /code/wait-for.sh /code/
COPY --from=builder /code/main .
COPY --from=builder /code/assets /code/assets
COPY --from=builder /code/html /code/html
CMD ["./main"]