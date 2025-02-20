FROM golang:1.23.6 as builder

WORKDIR /src

COPY . .

RUN go mod tidy
RUN go build -o /plugin-datasource main.go

# Use the official Grafana base image to serve the plugin
FROM grafana/grafana:latest

COPY --from=builder /plugin-datasource /etc/grafana/provisioning/datasources/

USER root