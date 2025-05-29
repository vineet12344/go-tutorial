FROM golang:1.24.3-alpine3.20
WORKDIR /app
COPY . /app
RUN go build -o dockerApi .
EXPOSE 4000
ENTRYPOINT [ "./dockerApi" ]