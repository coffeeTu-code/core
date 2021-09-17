#!/usr/bin/env bash

flatc --version

function gen() {
  dir=$1"_go"
  flatc -o $dir/ --grpc --go $1 && rsync -a $dir ../ && rm -r $dir

#  dir=$1"_py"
#  flatc -o $dir/ --python $1 && rsync -a $dir ../ && rm -r $dir
#
#  dir=$1"_cpp"
#  flatc -o $dir/ --grpc --cpp $1 && rsync -a $dir ../ && rm -r $dir
#
#  dir=$1"_rust"
#  flatc -o $dir/ --rust $1 && rsync -a $dir ../ && rm -r $dir
}

gen creative.fbs
