# Accept the Go version for the image to be set as a build argument.
ARG goVersion=1.13

# First stage: build the executable.
FROM golang:${goVersion}-alpine AS builder

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./cmd ./cmd
COPY ./generated ./generated
COPY ./pkg ./pkg
COPY ./main.go ./

RUN go build -o /bright main.go

# Final stage: the running container.
FROM alpine AS final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the first stage.
COPY --from=builder /bright /bright

# app files
COPY ./app.datestamp /app.datestamp
COPY ./app.timestamp /app.timestamp
COPY ./app.version /app.version

# static files
COPY ./templates /templates
COPY ./tpl /tpl
COPY ./static /static


# Declare the port on which the webserver will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8080

# Perform any further action as an unprivileged user.
USER nobody:nobody

# Run the compiled binary.
ENTRYPOINT ["/bright"]
CMD [  "servex", "--config", "/.bright.yaml"]