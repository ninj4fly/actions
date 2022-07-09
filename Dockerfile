FROM registry.access.redhat.com/ubi8/go-toolset AS builder

COPY app /app

USER root

RUN chmod a+w /app

USER 1001

WORKDIR /app

RUN go build main.go

FROM scratch

WORKDIR /

COPY --from=builder /app/main .

EXPOSE 80

ENTRYPOINT ["/main"]
