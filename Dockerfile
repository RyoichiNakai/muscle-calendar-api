# dev
# goのバイナリビルドをビルドするステージ
FROM golang:1.21 as dev

ARG SERVICE_NAME=muscle-calendar-api

WORKDIR /work/src/

COPY ./ ./
RUN go install github.com/cosmtrek/air@latest
RUN make mod build-linux

# release
# distroless: バイナリの実行に必要な依存のみが含まれているContainerImage
# パッケージマネージャやシェルも入っておらず、堅牢性が高くイメージサイズも小さい
# Goで生成したバイナリを実行するユースケースに適している
FROM gcr.io/distroless/static AS release

ARG RELEASE_VERSION
ARG COMMIT_HASH

ENV TZ="Asia/Tokyo"
ENV RELEASE_VERSION=${RELEASE_VERSION} \
    COMMIT_HASH=${COMMIT_HASH}

USER nonroot

COPY --chown=nonroot:nonroot --from=dev /work/src/build/app ./app

EXPOSE 8080

ENTRYPOINT ["./app"]
