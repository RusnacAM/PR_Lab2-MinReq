# Base go-image
FROM golang:1.17-alpine

# Create a directory for the app
RUN mkdir /app

# Copy all the files form the current dir to the app dir
COPY . /app

# Set working dir
WORKDIR /app

# GO build will build an executable file named client in curr dir
RUN go build -o producer .

# Expose port
EXPOSE 8000

# Run dining-hall
CMD [ "/app/producer" ]