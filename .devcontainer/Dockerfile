FROM golang:1.22

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y openjdk-17-jre && \
    apt-get autoremove -y && \
    apt-get clean -y && \
    rm -rf /var/lib/apt/lists/* && \
	rm -rf /var/cache/apt/*

RUN wget -O /tmp/questdb-7.4.2-no-jre-bin.tar.gz https://github.com/questdb/questdb/releases/download/7.4.2/questdb-7.4.2-no-jre-bin.tar.gz && \
	cd /opt && \
	tar xvzf /tmp/questdb-7.4.2-no-jre-bin.tar.gz && \
	rm /tmp/questdb-7.4.2-no-jre-bin.tar.gz && \
	ln -s questdb-7.4.2-no-jre-bin questdb

ENV JAVA_HOME=/usr/lib/jvm/java-17-openjdk-arm64

#EXPOSE 8812
EXPOSE 9000
#EXPOSE 9003
EXPOSE 9009

CMD [ "/bin/bash" ]
