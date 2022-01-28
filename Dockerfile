FROM debian:stretch-slim

WORKDIR /

COPY sample-k8s-scheduler /usr/local/bin

CMD ["sample-k8s-scheduler"]