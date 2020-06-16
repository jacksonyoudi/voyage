FROM  centos
RUN yum  -y install curl
ENTRYPOINT ["curl", "-s", "https://tool.lu/ip"]
ONBUILD RUN echo "hello world"