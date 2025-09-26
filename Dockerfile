# Use a lightweight base image
FROM alpine:latest

# Set environment variables
ENV SCRAPE_URL=""

# Set the working directory
WORKDIR /app

# Copy the built binary
COPY monster-scrape .

# Expose the port your Go app listens on
EXPOSE 8080
# Command to run the application
CMD ["./monster-scrape"]

ENTRYPOINT ["./monster-scrape"]