FROM golang:1.22

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y openjdk-17-jre sqlite3 unzip && \
    apt-get autoremove -y && \
    apt-get clean -y && \
    rm -rf /var/lib/apt/lists/* && \
	rm -rf /var/cache/apt/*

RUN wget -O /tmp/questdb-7.4.2-no-jre-bin.tar.gz https://github.com/questdb/questdb/releases/download/7.4.2/questdb-7.4.2-no-jre-bin.tar.gz && \
	cd /opt && \
	tar xvzf /tmp/questdb-7.4.2-no-jre-bin.tar.gz && \
	ln -s questdb-7.4.2-no-jre-bin questdb && \
	rm /tmp/questdb-7.4.2-no-jre-bin.tar.gz

ENV JAVA_HOME=/usr/lib/jvm/java-17-openjdk-arm64

RUN wget -O /tmp/duckdb_cli.zip https://github.com/duckdb/duckdb/releases/download/v0.10.2/duckdb_cli-linux-`uname -m`.zip && \
	mkdir -p /opt/duckdb && \
	cd /opt/duckdb && \
	unzip /tmp/duckdb_cli.zip && \
	ln -s /opt/duckdb/duckdb /bin/duckdb && \
	rm /tmp/duckdb_cli.zip

#EXPOSE 8812
EXPOSE 9000
#EXPOSE 9003
EXPOSE 9009

CMD [ "/bin/bash" ]
