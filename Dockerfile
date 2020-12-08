# Use the golang alpine image as the base stage of a multi-stage routine
FROM golang:1.15-alpine as base

# Set the working directory 
WORKDIR /blog

# Used for image scanning
FROM aquasec/trivy:0.14.0 as trivy

RUN trivy --debug --timeout 4m golang:1.15-alpine && \
  echo "No image vulnerabilities" > result

# Extend the base stage and create a new stage named dev
FROM base as dev

# Copy the go.mod and go.sum files to /goblog in the image's filesystem
COPY . .

# Install go module dependencies and verify in the image's filesystem
RUN go mod download
RUN go mod verify

# ENV sets an environment variable
ENV GOPATH /go
# Create GOPATH and PATH environment variables
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Print go environment for debugging purposes
RUN go env

# Build binary
# RUN go build -o main ./cmd/blog

# Install development dependencies to debug and live reload the server
RUN go get github.com/go-delve/delve/cmd/dlv \
  && go get github.com/githubnemo/CompileDaemon

# Provide meta data about the ports the container must expose
# port 8000 -> blog server port
# port 2345 -> debugger port
EXPOSE 8080 2345

# Extend the dev stage and create a new stage named test
FROM dev as test

# Copy the remaining server code into /goblog in the image's filesystem
COPY . .

# Disable CGO and run unit tests
RUN export CGO_ENABLED=0 && \
  go test -v ./...

# Extend thetest stage and create a new stage named build-stage
# FROM test as build-stage

# # Build the api with "-ldflags" aka linker flags to reduce the binary size
# # -s= disable symbol table
# # -w = disable DWARF code generation
# RUN GOOS=linux go build -ldflags "-s -w" -o main ./cmd/blog

# Extend the base stage and create a new stage named prod
# FROM base as prod

# # Copy only the files we want from a few stages into the prod stage
# COPY --from=trivy result secure
# COPY --from=build-stage /blog/main main

# # Create a new group and user, recursively change directory ownership, then allow the binary to be executed
# RUN addgroup parish && adduser -D -G parish parish \
#   && chown -R parish:parish ./ && \
#   chmod +x ./main

# # Change to a non-root user
# USER parish

# # Provide meta data about the port the container must expose
# EXPOSE 8080

# HEALTHCHECK CMD ["wget","-q","0.0.0.0:8080"]

# Provde the default command
CMD ["./main"]