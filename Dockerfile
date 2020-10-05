FROM archer.holygrail.com:5000/chromedriver
LABEL MAINTAINER=staroffish
RUN mkdir -p /workspace/
COPY ./ /workspace/
RUN chmod +x /workspace/am
ENTRYPOINT ["/workspace/am"]
WORKDIR /
