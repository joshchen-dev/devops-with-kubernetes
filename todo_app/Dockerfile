FROM golang:1.26.4-alpine3.24 AS build
WORKDIR /app
COPY . .
RUN go build -o /log_output/app .

FROM alpine:3.24.1
WORKDIR /app
COPY --from=build /log_output/app .
COPY static/ static/
ENTRYPOINT [ "./app" ]