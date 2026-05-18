FROM golang:1.26.3-alpine AS build
ENV CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go build \
    -o /tmp/helloworld \
    main.go

FROM scratch
COPY --from=build /tmp/helloworld ./helloworld
ENTRYPOINT [ "./helloworld" ]
EXPOSE 8080