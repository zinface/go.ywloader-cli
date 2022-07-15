package youwant

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"gitee.com/zinface/go.ywloader-cli/extenstions/logs"
	"gitee.com/zinface/go.ywloader-cli/models"
	"gitee.com/zinface/go.ywloader-cli/utils"
)

// 此处开始实现 add 命令

var alog = &logs.Logs{
	Prefix: "add",
}

// 定义 命令 step
const (
	addYouwantItemLabel         = "> 1.设置条目标题(label): "
	addYouwantItemDetail        = "> 2.设置条目简介(detail): "
	addYouwantItemTypes         = "> 3.设置条目类别(格式为:a|b|c): "
	addYouwantItemAction        = "> 4.设置模板动作(默认:shell): "
	addYouwantItemSehllDelay    = "> 5.设置动作延时(默认:0): "
	addYouwantItemShellCommand  = "> 6.设置模板指令(设置第%v条指令): "
	addYouwantItemVSCodeCommand = "> 7.设置VSCode指令(1.刷新资源管理器,2.重新加载窗口): "
)

// addYouwantItem 获取一项内容
// 走完一遍流程
func addYouwantItem() models.Youwant {
	alog.Println("准备创建模板")
	var _t_label string
	var _t_detail string
	var _t_types string
	var _t_action string
	var _t_shell_delay int
	var _t_shell_commands = []string{}
	var _t_vscommands = []string{}

	_t_label = utils.GetStdinStringValue(addYouwantItemLabel, "")
	_t_detail = utils.GetStdinStringValue(addYouwantItemDetail, "")
	_t_types = utils.GetStdinStringValue(addYouwantItemTypes, "")
	_t_action = utils.GetStdinStringValue(addYouwantItemAction, "shell")
	_t_shell_delay = utils.GetStdinNumberValue(addYouwantItemSehllDelay, 0)

	var _command_cnt = 1

	for {
		var _command = utils.GetStdinStringValue(fmt.Sprintf(addYouwantItemShellCommand, _command_cnt), "")
		if len(_command) == 0 {
			break
		}
		_t_shell_commands = append(_t_shell_commands, _command)
		_command_cnt += 1
	}

	var _vstype = utils.GetStdinNumberValue(addYouwantItemVSCodeCommand, 1)
	switch _vstype {
	case 1:
		_t_vscommands = append(_t_vscommands, models.RefreshFilesExplorer)
	case 2:
		_t_vscommands = append(_t_vscommands, models.ReloadWindow)
	}

	var want = models.Youwant{
		Label:  _t_label,
		Detail: _t_detail,
		Type:   _t_types,
		Template: models.Template{
			Action: _t_action,
			Shell: models.Shell{
				Delay:    _t_shell_delay,
				Commands: _t_shell_commands,
			},
			VSCommands: _t_vscommands,
		},
	}

	return want
}

func readDir(dirPath string) []string {
	var filelist []string
	flist, e := ioutil.ReadDir(dirPath)
	if e != nil {
		fmt.Println("read file error")
		return filelist
	}
	for _, fpath := range flist {
		if fpath.IsDir() {
			filelist = append(filelist, readDir(dirPath+"/"+fpath.Name())...)
		} else {
			filelist = append(filelist, filepath.Join(dirPath, fpath.Name()))
		}
	}
	return filelist
}

// addFromCommandFileFlags 尝试从 --file 参数中获取指定的文件内容
func addFromCommandFileFlags(cmd *cobra.Command, args []string) models.FileItems {
	var files = models.FileItems{}
	// 获取所有文件路经
	var fpaths []string
	// for {
	fpaths, err := cmd.Flags().GetStringArray("file")
	if err != nil {
		alog.Println(err.Error())
		return models.FileItems{}
	}

	// 处理剩下参数，可能包含文件
	for _, path := range args {
		if utils.IsFile(path) {
			// 如果是文件
			// alog.Printf("处理文件: %s\n", path)
			fpaths = append(fpaths, path)
		} else if utils.IsDir(path) {
			// 如果是目录
			// alog.Printf("处理目录: %s\n", path)
			fpaths = append(fpaths, readDir(path)...)
		}
	}

	// 路径去重
	keys := make(map[string]string, 0)
	var newFpaths = []string{}
	for _, fpath := range fpaths {
		if _, ok := keys[fpath]; !ok {
			keys[fpath] = ""
			newFpaths = append(newFpaths, fpath)
		}
	}
	for _, fpath := range newFpaths {
		if utils.FileExists(fpath) {
			data, err := ioutil.ReadFile(fpath)
			if err != nil {
				panic(err)
			}
			base64Str := base64.StdEncoding.EncodeToString(data)

			finfo, err := os.Stat(fpath)
			var perm int32 = 0
			if err == nil {
				perm = int32(finfo.Mode().Perm())
			}

			files = append(files,
				models.FileItem{
					Name:       fpath,
					Base64:     base64Str,
					Permission: perm,
				})
			log.Println("添加文件到项目:", fpath)
		}
	}
	return files
}

// addHandler 指令实现
func AddHandler(cmd *cobra.Command, args []string) {

	// 使用的配置文件
	var useConfigFile = useConfigFilePath(cmd)
	alog.UseConfig(useConfigFile)

	// 文件内容
	var wants models.Youwants

	// 存在该文件将尝试读取
	if utils.FileExists(useConfigFile) {
		data, err := os.ReadFile(useConfigFile)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &wants)
		if err != nil {
			panic(err)
		}
	} else {
		// 否则将询问是否创建
		for {
			var _create = utils.GetStdinStringValue("配置文件不存在是否创建?(yes/no): ", "")
			if strings.Contains(_create, "yes") {
				// 返回 yes 将表示创建该文件
				if !initGenerateYwloaderFile(useConfigFile) {
					os.Exit(0)
				}
				break
			}
			if strings.Contains(_create, "no") {
				alog.Println("已放弃创建.")
				os.Exit(0)
			}
		}
	}

	// 添加一项配置
	var want = addYouwantItem()
	want.Template.Files = addFromCommandFileFlags(cmd, args)

	// 配置验证
	if !want.IsValid() {
		for {
			var _continue = utils.GetStdinStringValue("缺失关键属性值(Label), 是否继续?(yes/no): ", "")
			if strings.Contains(_continue, "yes") {
				break
			}
			if strings.Contains(_continue, "no") {
				alog.Println("已放弃创建")
				os.Exit(0)
			}
		}
	}

	wants = append(wants, want)
	err := wants.SaveFile(useConfigFile)
	if err != nil {
		alog.Println("项目创建失败.")
		panic(err)
	}

	alog.Println("项目创建成功.")
	fmt.Printf("%v\n", want.ToJsonString())
}
