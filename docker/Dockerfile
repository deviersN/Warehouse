# WAREHOUSE DOCKERFILE with golang:latest image

FROM golang:alpine

RUN mkdir /resources

WORKDIR /go/src/github.com/deviersN/Warehouse

COPY ./ /go/src/github.com/deviersN/Warehouse

RUN rm go.mod

ENV PATH=${PATH}:/go/src/github.com/deviersN/Warehouse

# Check and download mdules
RUN go mod init github.com/deviersN/Warehouse
RUN go mod download
RUN go mod verify

RUN go get -d -v ./...

# Install binary
RUN go install -v ./...

EXPOSE 8888

# Run executable
CMD ["Warehouse", "/resources/instructions.txt"]
