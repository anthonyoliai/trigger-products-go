FROM golang:1.23

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod tidy

# Fetch the additional package
RUN go get github.com/anthonyoliai/trigger-products-go/storage

# Copy the source code. Note the slash at the end
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /triggers-products

CMD ["/triggers-products"]
