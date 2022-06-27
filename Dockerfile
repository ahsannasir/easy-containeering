
# Using the golang version as my machine
FROM golang:1.18-alpine
# Set working directory to App
WORKDIR /app
# Copy project code to container in /app directory
COPY . .
# Build the project and generate an executable
RUN go build -o /mlcicd
# API runs on port 5433
EXPOSE 5433
# Start the executable
CMD [ "/mlcicd" ]