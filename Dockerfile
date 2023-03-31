FROM golang:1.19.7-alpine3.17 as builder
RUN apk add --no-cache gcc git make musl-dev

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
COPY ./Makefile /app/Makefile

COPY ./bindings /app/bindings
COPY ./components /app/components
COPY ./utils /app/utils

COPY ./.git /app/.git
WORKDIR /app
RUN make build

# Node
FROM alpine:3.17 as kanvas-node
COPY --from=builder /app/bin/kanvas-node /usr/local/bin
ENTRYPOINT ["kanvas-node"]

# Stateviz
FROM alpine:3.17 as kanvas-stateviz
COPY --from=builder /app/bin/kanvas-stateviz /usr/local/bin
CMD ["kanvas-stateviz"]

# Batcher
FROM alpine:3.17 as kanvas-batcher
COPY --from=builder /app/bin/kanvas-batcher /usr/local/bin
ENTRYPOINT ["kanvas-batcher"]

# Validator
FROM alpine:3.17 as kanvas-validator
COPY --from=builder /app/bin/kanvas-validator /usr/local/bin
ENTRYPOINT ["kanvas-validator"]
