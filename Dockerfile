# syntax=docker/dockerfile:1

FROM alpine:3.18

# Set destination for COPY
WORKDIR /.

# Download Go modules
COPY main ./


# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 4880

# Run
CMD ["/main"]