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

# 一些文章

[深入浅出 FlatBuffers 之 Schema ++++++++++++ 一缕殇流化隐半边冰霜](https://halfrost.com/flatbuffers_schema/#toc-14)