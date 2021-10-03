#!/usr/bin/env bash

git submodule update --init
sudo apt install autoconf libtool
cd fdkaac-lib
bash autogen.sh
./configure --prefix=`pwd`/../fdkaac-lib-objs
make
make install
cd -
