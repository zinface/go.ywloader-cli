# 默认将构建 ywloader
all:
	go build cli/ywloader.go

# make install 将会安装 ywloader 到 $GOPATH/bin 中
install:
	go install cli/ywloader.go

# 已安装 shellcheck 将可使用
SHELLCHECK=$(shell which shellcheck)
ifneq (${SHELLCHECK}, "")
shellcheck:
	shellcheck disable=SC2016,SC2119,SC2155,SC2206,SC2207 cli/bash-completions/ywloader 
endif

# 已安装 pkg-config 并且已安装 bash-completion 将可使用
PKGCONFIG=$(shell which pkg-config)
ifneq (${PKGCONFIG}, "")
BASH_COMPATDIR=$(shell pkg-config --variable=compatdir bash-completion)
ifneq (${BASH_COMPATDIR}, "")
bash-completions:
	@echo "Please run: source cli/bash-completions/ywloader"

install-bash-completions:
	@echo "cli/bash-completions/ywloader will be installed ywloader in ${BASH_COMPATDIR}"
	cp cli/bash-completions/ywloader ${BASH_COMPATDIR}
endif
endif