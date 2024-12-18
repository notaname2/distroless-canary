ARG GO_VERSION=1.23

 
# STAGE 1: building the executable
FROM golang:${GO_VERSION}-alpine AS build
 
RUN apk add --no-cache git


WORKDIR /apps
COPY ./apps/ ./
COPY ./go.mod ./

RUN pwd
RUN ls -la

RUN CGO_ENABLED=0 go build \
	-installsuffix 'static' \
	-o /server ./server

RUN CGO_ENABLED=0 go build \
	-installsuffix 'static' \
	-o /time ./mytime

RUN CGO_ENABLED=0 go build \
	-installsuffix 'static' \
	-o /hello ./hello

RUN CGO_ENABLED=0 go build \
	-installsuffix 'static' \
	-o /canary ./canary

# STAGE 2: build the container to run
FROM gcr.io/distroless/static AS final
 
USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /server /app/server
COPY --from=build --chown=nonroot:nonroot /time /app/bin/time
COPY --from=build --chown=nonroot:nonroot /hello /app/bin/hello

# Canaries
COPY --from=build --chown=nonroot:nonroot /canary /bin/sh




EXPOSE 8080:8080/tcp

ENTRYPOINT ["/app/server"]


