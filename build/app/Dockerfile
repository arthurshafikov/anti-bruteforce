FROM golang:1.16.2 as build

ENV BIN_FILE /opt/app/app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS

RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/*

FROM alpine:3.9

LABEL MAINTAINERS="thewolf27"

ENV BIN_FILE "/opt/app/app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CODE_DIR /go/src/

ENV CONFIG_FOLDER "/opt/app/configs"
COPY --from=build ${CODE_DIR}/configs ${CONFIG_FOLDER}

# Add docker-compose-wait tool -------------------
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.2/wait /wait
RUN chmod +x /wait

CMD /wait && \
    ${BIN_FILE} --configFolder ${CONFIG_FOLDER}
