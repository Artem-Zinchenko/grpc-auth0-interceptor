FROM golang:1.19 as builder

WORKDIR /code

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# buid
COPY . ./
RUN go build -v -o  /app cmd/catalog/main.go

FROM gcr.io/distroless/base
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single

COPY --from=builder /app /app
EXPOSE 8080
ENTRYPOINT ["./app"]


