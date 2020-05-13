FROM arm64v8/golang:1.14-alpine as builder

RUN mkdir -p /podinfo/

WORKDIR /podinfo

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" \
    -a -o bin/podinfo cmd/podinfo/*

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" \
    -a -o bin/podcli cmd/podcli/*

FROM arm64v8/alpine:3.11

RUN addgroup -S app \
    && adduser -S -g app app \
    && apk --no-cache add \
    curl openssl netcat-openbsd

WORKDIR /home/app

COPY --from=builder /podinfo/bin/podinfo .
COPY --from=builder /podinfo/bin/podcli /usr/local/bin/podcli
COPY ./ui ./ui
RUN chown -R app:app ./

USER app

CMD ["./podinfo"]
