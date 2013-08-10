# Container running Eventstore
#
# VERSION               0.1
FROM ubuntu
MAINTAINER Pieter Joost van de Sande "pj@born2code.net"

# make sure the package repository is up to date
RUN apt-get update

# install packages required to build mono and the eventstore
RUN apt-get install -y git git-core subversion
RUN apt-get install -y autoconf automake libtool gettext libglib2.0-dev libfontconfig1-dev mono-gmcs
RUN apt-get install -y build-essential

# get eventstore and build it
RUN git clone https://github.com/EventStore/EventStore.git /var/local/EventStore --depth=1

# install patched mono
RUN (cd /var/local/EventStore; ./src/EventStore/Scripts/get-mono-patched.sh)

# set env vars for mono
ENV PATH /usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/opt/mono/bin
ENV LD_LIBRARY_PATH /opt/mono/lib/
ENV PKG_CONFIG_PATH /opt/mono/lib/pkgconfig

# build event store
RUN (cd /var/local/EventStore; ./build.sh full configuration=release)

# setup directory structure
ENV EVENTSTORE_BLD /var/local/EventStore/bin/eventstore/release/anycpu
ENV EVENTSTORE_ROOT /opt/eventstore
ENV EVENTSTORE_BIN /opt/eventstore/bin
ENV EVENTSTORE_DB /opt/eventstore/db
ENV EVENTSTORE_LOG /opt/eventstore/logs

# create directory structure
RUN mkdir -p $EVENTSTORE_ROOT
RUN mkdir -p $EVENTSTORE_DB
RUN mkdir -p $EVENTSTORE_LOG

# cleanup
RUN mv "$EVENTSTORE_BLD" "$EVENTSTORE_BIN"
RUN rm -rf "/var/local/EventStore"

# remove packages to reduce container size
RUN apt-get remove --purge -y git git-core subversion
RUN apt-get remove --purge -y autoconf automake libtool gettext libglib2.0-dev libfontconfig1-dev mono-gmcs
RUN apt-get remove --purge -y build-essential

# remove mono
RUN rm -rf /mono

# export the http and tcp port
EXPOSE 2113:2113
EXPOSE 1113:1113

# set entry point to eventstore executable
ENTRYPOINT ["mono-sgen", "/opt/eventstore/bin/EventStore.SingleNode.exe", "--log=/opt/eventstore/logs", "--db=/opt/eventstore/db"]

# bind it to all interfaces by default
CMD ["--ip=127.0.0.1", "--http-prefix=http://*:2113/"]
