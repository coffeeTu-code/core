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

# cat std_rank.proto| tr 'A-Z' 'a-z'
flatc -I . --proto std_rank.proto && gen std_rank.fbs

