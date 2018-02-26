FROM golang:1.9.4 as build

# Install golang/dep required to build the project
# We need to run `docker build --no-cache .` to update those dependencies
RUN go get github.com/golang/dep/cmd/dep

# Gopkg.toml and Gopkg.lock lists project dependencies
# These layers are only re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/github.com/benkim0414/superego/
WORKDIR /go/src/github.com/benkim0414/superego/
# Install library dependencies
RUN dep ensure -vendor-only

# Copy all project and build it
# This layer is rebuilt when ever a file has changed in the project directory
COPY . /go/src/github.com/benkim0414/superego/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/superego ./cmd/superego

# This results in a single layer image
FROM alpine:latest
RUN apk add --no-cache ca-certificates

# Copy executable and application crednetials
COPY --from=build /bin/superego /go/src/github.com/benkim0414/superego/superego-b60ee40757c8.json /bin/

# Set envrionment variables to provide credentials to the application
ENV GCP_PROJECT_ID superego-173310
ENV GOOGLE_APPLICATION_CREDENTIALS /bin/superego-b60ee40757c8.json

# Expose ports and run the application
EXPOSE 8080 8081
CMD ["/bin/superego"]
