FROM alpine:3.4

ENV PROJECT_NAME=httptest
ARG DUMB_INIT_VERSION=1.2.0
ARG DOCKERIZE_VERSION=0.2.0

RUN apk update && apk add --no-cache bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/Yelp/dumb-init/releases/download/v${DUMB_INIT_VERSION}/dumb-init_${DUMB_INIT_VERSION}_amd64 /usr/local/bin/dumb-init
RUN chmod +x /usr/local/bin/dumb-init

ADD https://github.com/jwilder/dockerize/releases/download/v${DOCKERIZE_VERSION}/dockerize-linux-amd64-v${DOCKERIZE_VERSION}.tar.gz /tmp
RUN tar -C /usr/local/bin -xzvf /tmp/dockerize-linux-amd64-v${DOCKERIZE_VERSION}.tar.gz && rm /tmp/dockerize-linux-amd64-v${DOCKERIZE_VERSION}.tar.gz
RUN chown root /usr/local/bin/dockerize && chgrp root /usr/local/bin/dockerize && chmod +x /usr/local/bin/dockerize

COPY ./bin/linux/${PROJECT_NAME} /bin/

RUN /usr/sbin/addgroup go && \
    /usr/sbin/adduser -D -G go go

USER go

ENTRYPOINT ["/usr/local/bin/dumb-init", "--"]

#######################################################################################################

CMD ["/bin/httptest"]
