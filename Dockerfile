FROM debian:bullseye-slim as build

RUN apt-get update && apt-get -y install build-essential autoconf libtool

COPY /fdkaac-lib /fdkaac-lib
WORKDIR /fdkaac-lib

RUN ./autogen.sh
RUN ./configure --prefix=/fdkaac-objs
RUN make
RUN make install

FROM scratch as artifacts

COPY --from=build /fdkaac-objs /artifacts
WORKDIR /artifacts
