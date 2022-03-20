FROM scratch
LABEL maintainer="maintainer@fluffy-bunny.com"
LABEL description="starter kit"

WORKDIR /app
COPY cli .

ENV PATH=/app:${PATH}

ENTRYPOINT ["/app/cli"]