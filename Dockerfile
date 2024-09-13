# Builder stage
FROM golang:1.22.5-bullseye AS builder

# Install dependencies
# RUN apt-get update && apt-get install -y wkhtmltopdf

# Set the working directory
WORKDIR /app
ENV CGO_ENABLED=0

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application using make
RUN make build

# Debug: List contents of /app after build
RUN ls -l /app

# Final stage
FROM debian:bullseye-slim

WORKDIR /app

# Make dist directory
RUN mkdir -p /dist

# Install wkhtmltopdf and its dependencies
RUN apt-get update && apt-get install -y wkhtmltopdf

# Verify installation of wkhtmltopdf
RUN which wkhtmltopdf

# Create a non-root user
RUN addgroup --system appgroup && adduser --system --ingroup appgroup appuser

# Copy built artifacts from the builder stage
COPY --from=builder /app/shipment-service-shipping-label /app/
COPY --from=builder /app/dist /app/dist/

# Set ownership
RUN chown appuser:appgroup /app/shipment-service-shipping-label
RUN chown -R appuser:appgroup /dist

EXPOSE 8003

# Run as non-root user
USER appuser

CMD ["/app/shipment-service-shipping-label"]
