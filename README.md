
# 使用go语言和chain33链上部署的solidity合约进行交互

## 开发
 1. 使用 solidity 实现合约 
 2. 开发工具可以选择  
      * 熟悉go : vscode & solidity & go
      * 熟悉js ： vscode & solidity & js & hardhat 
 3. go: abigen & solc  (编译) 
 4. 应用： 直接调用 abigen 生成的go代码， 部署合约及和合约交互。
 
### solidity 合约

资料
  1. https://docs.soliditylang.org/en/latest/ 
  2.  https://learnblockchain.cn/docs/solidity/index.html

### vscode & solidity & go (开发)

 1. 插件 solidity,  solidity-extended, solidity-debugger
 2. 配置和版本切换 https://marketplace.visualstudio.com/items?itemName=JuanBlanco.solidity

```
   // "solidity.compileUsingRemoteVersion": "https://github.com/ethereum/solc-bin/tree/gh-pages/bin",
    "solidity.compileUsingRemoteVersion": "v0.8.17+commit.8df45f5fa",
    "solidity.compileUsingLocalVersion": "/home/linj/work/src/github.com/linj-disanbo/solidity-chain33/soljson-v0.8.17+commit.8df45f5f.js",
    "solidity.defaultCompiler": "localFile",
```

### abigen & solc  (编译)

编译和代码生成

```
./tools/abigen --solc ./tools/solc  --sol user.sol  --pkg user  --out user.go
go build cmd/main.go
```

其中 
 1. solc: Version: 0.8.0+commit.c7dfd78e.Linux.g++


##  测试环境


### remix 模拟环境 

http://remix.ethereum.org/


### chain33 环境

```
 plugin 版本 : git@github.com:33cn/plugin.git f3c6aa73066e872c6f7e3d3cd59f1a9119ba7315 (当前master)
  固定版本为了不用更新配置
```

 1. 将 chain33-sole-evm 下文件复制到 plugin 仓库下
 2. docker build -t chain33-evm-user .
 3. docker-compose up -d
 4. docker  cp main  soliditychain33_chain33-evm_1:/app
 5. docker exec soliditychain33_chain33-evm_1 /app/main

 main 为测试程序

#### 钱币初始化

solo 不收手续费， 不一定需要

```
docker exec soliditychain33_chain33-evm_1 ./chain33-cli seed save -p 123123qweqwe -s "remind evolve collect figure road rather dynamic later praise earn alter tool unlock cousin web"
docker exec soliditychain33_chain33-evm_1 ./chain33-cli wallet unlock -p  123123qweqwe
docker exec soliditychain33_chain33-evm_1 ./chain33-cli  account  import_key  -t 2 -l miner -k 0x1130d8e8b921cc1815048f1ef69a1b9cf1b2d8283f45aa8100b55b4bb6510fee
```


### remix

 vscode 和 浏览器 remix 打通， remix 插件

也可以直接用 remix 


## local chain

其他环境

http://124.71.110.109:8547
