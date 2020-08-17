# We specify the base image we need for our
# go application - Alpine image didn't have git or gcc. 
FROM golang:1.14-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git make build-base

# We create an /app directory within our
# image that will hold our application source
# files
RUN mkdir /cpx
# We copy everything in the root directory
# into our /app directory
ADD . /cpx
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /cpx
# Add this go mod download command to pull in any dependencies
RUN go mod download
# Run the CPX tests in a docker environment
RUN go test -v .

