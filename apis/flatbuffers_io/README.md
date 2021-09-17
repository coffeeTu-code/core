# FlatBuffers在Mac OS的安装步骤

```shell
# 需要安装cmake
brew install cmake

# 1.下载源码
git clone https:# github.com/google/flatbuffers.git

# 2.生产makefile文件
cd flatbuffers
cmake -G "Unix Makefiles"

# 3.安装
make
make install

# 使用命令
flatc --version
# flatc version 2.0.0

```