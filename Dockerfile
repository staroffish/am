FROM busybox:glibc
LABEL MAINTAINER=staroffish
RUN mkdir -p /workspace/
COPY ./ /workspace/
RUN chmod +x /workspace/am
ENTRYPOINT ["/workspace/am"]
WORKDIR /
