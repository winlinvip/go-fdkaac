#!/usr/bin/env bash

git submodule update --init
cd fdkaac-lib
bash autogen.sh
./configure --prefix=`pwd`/../fdkaac-lib-objs
make
make install
cd -
