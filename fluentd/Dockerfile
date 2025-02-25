# fluentd/Dockerfile

FROM fluent/fluentd:v1.17.1-debian-1.1
USER root
RUN ["gem", "install", "fluent-plugin-opensearch", "--no-document"]
USER fluent