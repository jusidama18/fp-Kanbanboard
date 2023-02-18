FROM golang:1.18-alpine as builder

RUN mkdir /app
WORKDIR /app

COPY . /app

RUN go build -o Kanbanboard ./main.go

# --- #
FROM alpine:latest 

RUN mkdir /app

WORKDIR /app
COPY --from=builder /app /app

ENTRYPOINT [ "./Kanbanboard" ]