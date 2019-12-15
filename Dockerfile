# Build Stage
FROM martient/ToDefine-build:1.13 AS build-stage

LABEL app="build-ToDefine"
LABEL REPO="https://github.com/Martient/ToDefine"

ENV PROJPATH=/go/src/github.com/Martient/ToDefine

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/Martient/ToDefine
WORKDIR /go/src/github.com/Martient/ToDefine

RUN make build-alpine

# Final Stage
FROM martient/ToDefine

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/Martient/ToDefine"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/ToDefine/bin

WORKDIR /opt/ToDefine/bin

COPY --from=build-stage /go/src/github.com/Martient/ToDefine/bin/ToDefine /opt/ToDefine/bin/
RUN chmod +x /opt/ToDefine/bin/ToDefine

# Create appuser
RUN adduser -D -g '' ToDefine
USER ToDefine

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/ToDefine/bin/ToDefine"]
