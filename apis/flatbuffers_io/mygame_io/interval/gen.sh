#!/usr/bin/env bash

flatc --version

function gen() {
  godir=$1"_go"
  flatc -o $godir/ --go $1 && rsync -a $godir ../ && rm -r $godir

  pydir=$1"_py"
  flatc -o $pydir/ --python $1 && rsync -a $pydir ../ && rm -r $pydir

  cppdir=$1"_cpp"
  flatc -o $cppdir/ --cpp $1 && rsync -a $cppdir ../ && rm -r $cppdir
}

gen mygame.fbs
