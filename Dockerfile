# Build stage (compiles the Go binary)
FROM golang:1.23.5 as builder

WORKDIR /src

COPY . .  

RUN go mod tidy  # Ensure dependencies are correct
RUN go build -o /plugin-datasource main.go  # Compile the Go binary

# Final Grafana image
FROM grafana/grafana:latest

# Create a directory for the plugin
RUN mkdir -p /var/lib/grafana/plugins/my-plugin

# Copy the Go binary from the builder stage
COPY --from=builder /plugin-datasource /var/lib/grafana/plugins/my-plugin/plugin-datasource

# Copy metadata file (required for Grafana to detect the plugin)
COPY plugin.json /var/lib/grafana/plugins/my-plugin/

CMD ["/run.sh"]