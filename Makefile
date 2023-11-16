# 默认将构建 ywloader
all:
	@test -f extra/bash-completion/README.md || git submodule update --init --recursive
	@go build

# make install 将会安装 ywloader 到 $GOPATH/bin 中
install:
	go install

# 已安装 shellcheck 将可使用
SHELLCHECK=$(shell which shellcheck)
ifneq (${SHELLCHECK}, "")
shellcheck:
	shellcheck disable=SC2016,SC2119,SC2155,SC2206,SC2207 extra/bash-completion/ywloader 
endif

# 已安装 pkg-config 并且已安装 bash-completion 将可使用
PKGCONFIG=$(shell which pkg-config)
ifneq (${PKGCONFIG}, "")
BASH_COMPATDIR=$(shell pkg-config --variable=compatdir bash-completion)
ifneq (${BASH_COMPATDIR}, "")
bash-completions:
	@echo "Please run: source extra/bash-completion/ywloader"

install-bash-completions:
	@echo "extra/bash-completion/ywloader will be installed ywloader in ${BASH_COMPATDIR}"
	cp extra/bash-completion/ywloader ${BASH_COMPATDIR}
endif
endif

# 进入调试模式，并指定调试文件为当前目录下的 ywloader-debug.file
enter-debug-mode:
	cd extra/bash-completion && export BASH_COMP_DEBUG_FILE=`pwd`/ywloader-debug.file && bash