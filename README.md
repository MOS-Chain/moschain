# About

This depository is a fork of xuperchain. Add transfer fee and mortgage reward.

For the original [moschain](https://github.com/MOS-Chain/moschain) depository, please click the link.

[中文说明](#中文说明-1)
-----
## What is MOSChain

**MOSChain**, a highly flexible blockchain architecture based on XuperChain.
 
MOSChain is the underlying solution for union networks with following highlight features:

* **High Performance**

    * Creative XuperModel technology makes contract execution and verification run parallelly.
    * TDPoS ensures quick consensus in a large scale network.
    * WASM VM using AOT technology.

* **Solid Security**

    * Contract account protected by multiple private keys ensures assets safety.
    * Flexible authorization system supports weight threshold, AK sets and could be easily extended.

* **High Scalability**

    * Robust P2P network supports a large scale network with thousands of nodes.
    * Branch management on ledger makes automatic convergence consistency and supports global deployment.

* **Multi-Language Support**: Support pluggable multi-language contract VM using XuperBridge technology.

* **Flexibility**:  Modular and pluggable design provides high flexibility for users to build their blockchain solutions for various business scenarios.

## Quick start

### Requirements

* OS Support: Linux and Mac OS
* Go 1.12.x or later
* GCC 4.8.x or later
* Git

### Build

Clone the repository

```
git clone https://github.com/MOS-Chain/moschain
```

Enter the moschain folder and build the code:

```
cd moschain
make
```

Note that if you are using Go 1.11 or later, go modules are used to download 3rd-party dependencies by default. You can also disable go modules and use the prepared dependencies under vendor folder.

Run test:
```
make test
```

Use Docker to build moschain see [docker build](core/scripts/README.md)

### Run 

There is an output folder if build successfully. Enter the output folder, create a default chain firstly:

```
cd ./output
./xchain-cli createChain
```

By doing this, a blockchain named "xuper" is created, you can find the data of this blockchain at `./data/blockchain/xuper/`.

Then start the node and run MOSChain full node servers:

```
nohup ./xchain &
```

By default, the `xuper` chain will produce a block every 3 seconds, try the following command to see the `trunkHeight` of chain and make sure it's growing.

```
./xchain-cli status
```

## Run with Docker

### Build image

```bash
docker build . -t moschain
```

### Run image

```bash
# run xchain daemon
docker run -d -p 37101:37101 -p 47101:47101 --rm --name mchian moschain
# enter running container
docker exec -ti mchain bash
# run command
./xchain-cli status
```

> This is only a demo for local single container, you can use volume to mount and overwrite configurations.

## License

MOSChain is under the [Apache License, Version 2.0](https://github.com/xuperchain/xuperchain/blob/master/LICENSE).


=====

# 中文说明

## MOSChain?

**MOSChain**是基于百度超级链的构建超级联盟网络的底层方案。

核心特点

* **高性能**
    * 百度原创的XuperModel模型，真正实现了智能合约的并发执行和验证。
    * TDPOS算法确保大规模节点下的快速共识。
    * 使用AOT加速的WASM虚拟机，合约运行速度接近native程序。

* **更安全**
    * 多私钥保护的账户体系。
    * 鉴权支持权重累计、集合运算等灵活的策略。

* **易扩展**
    * 鲁棒的P2P网络，支持广域网超大规模节点。
    * 底层账本支持分叉管理，自动收敛一致性，实现真正全球化部署。

* **多语言开发智能合约**
    * 百度原创的XuperBridge技术，可插拔多语言虚拟机。

* **高灵活性**
    * 可插拔、插件化的设计使得用户可以方便选择适合自己业务场景的解决方案。

## 快速试用

### 环境配置

* 操作系统：支持Linux以及Mac OS
* 开发语言：Go 1.12.x及以上
* 编译器：GCC 4.8.x及以上
* 版本控制工具：Git

### 构建

克隆MOSChain仓库
```
git clone https://github.com/MOS-Chain/moschain
```

编译
```
cd moschain
make
```

跑单测
```
make test
```

使用docker来编译moschain见[docker build](core/scripts/README.md)

单机版xchain
```
cd ./output
./xchain-cli createChain
nohup ./xchain &
./xchain-cli status
```

## 容器运行

### 编译镜像

```bash
docker build . -t moschain

### 运行镜像

```bash
# 运行容器 daemon
docker run -d -p 37101:37101 -p 47101:47101 --rm --name mchain moschain
# 进入容器
docker exec -ti mchain bash
# 运行指令
./xchain-cli status
```

> 本地容器化运行的示例，实际场景中可以用卷的方式挂载并覆盖配置。


## 许可证
MOSChain使用的许可证是Apache 2.0


