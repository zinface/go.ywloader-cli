# ywloader - Youwant 命令行工具

```
ywloader 可执行程序

Usage:
  ywloader [command]

Available Commands:
  add         添加项目到仓库索引
  cat         查看指定项目的文件
  change      修改一条项目
  completion  bash补全脚本
  del         删除一条项目
  help        Help about any command
  init        初始化当前目录为 Youwant 仓库
  list        列出仓库项目索引
  logo        打印可执行程序 logo 头部
  remote      远程仓库命令
  search      搜索仓库项目索引
  serve       启动一个web服务(默认:8080)
  show        显示项目详细信息
  update      更新一条项目的文件列表
  use         使用快速项目模板
  version     打印版本信息

Flags:
  -h, --help   help for ywloader

Use "ywloader [command] --help" for more information about a command.
```

## ywloader-cli 命令的安装与补全功能

- 安装
  ```sh
  git clone --recurse-submodules https://gitee.com/zinface/ywloader-cli
  cd ywloader-cli
  
  # 项目中包含 Makefile
  make install    # 实际执行 go install
  ```

- bash-completion 脚本(正在重构中...)
  ```sh
  # 基于 pkg-config 命令
  # 获取 bash-completion 包的 compatdir 变量得到可安装的位置
  sudo make install-bash-completions
  ```