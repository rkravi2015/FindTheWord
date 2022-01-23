# Start with golang base image
FROM golang:1.16-alpine

# Add Maintainer Info
LABEL maintainer="Ravi Kumar <ravikr.1502015@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /FindTheWord

# Copy the source from the current directory to the Working Directory inside the container
COPY * ./

# Build the FindTheWord app
RUN go build main.go

# Run FindTheWord app
RUN ./main input.txt

# Command to display result
CMD [ "cat", "response.txt" ]