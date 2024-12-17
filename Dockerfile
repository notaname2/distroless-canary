ARG GO_VERSION 1.23

 
# STAGE 1: building the executable
FROM golang:${GO_VERSION}-alpine AS build
 
RUN apk add --no-cache git


COPY ./src/ ./
WORKDIR /src
COPY ./go.mod ./

RUN CGO_ENABLED=0 go build \
	-installsuffix 'static' \
	-o /server ./apps/server


# STAGE 2: build the container to run
FROM gcr.io/distroless/static AS final
 
USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /server /app/server

ENTRYPOINT ["/app/server"]


