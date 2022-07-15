package logs

import "log"

// 日志打印说明
const (
	// 配置文件
	_logConfigLocation       = "配置文件位置:"
	_logConfigNotExist       = "配置文件不存在:"
	_logConfigIsExist        = "配置文件已存在:"
	_logConfigCreated        = "配置文件已创建:"
	_logConfigCreatedFail    = "配置文件创建失败:"
	_logConfigCreatedSuccess = "配置文件创建成功:"
	_logConfigCreatedDiscard = "配置文件已放弃创建:"

	// 用户配置文件
	_logUserConfigLocation       = "用户配置文件位置:"
	_logUserConfigNotExist       = "用户配置文件不存在:"
	_logUserConfigIsExist        = "用户配置文件已存在:"
	_logUserConfigCreated        = "用户配置文件已创建:"
	_logUserConfigCreatedFail    = "用户配置文件创建失败:"
	_logUserConfigCreatedSuccess = "用户配置文件创建成功:"
	_logUserConfigCreatedDiscard = "用户配置文件已放弃创建:"

	// 文件
	_logFileCreated        = "文件已创建:"
	_logFileIsExist        = "文件已存在:"
	_logFileNotExits       = "文件不存在:"
	_logFileCreatedFail    = "文件创建失败:"
	_logFileCreatedSuccess = "文件创建成功:"
	_logFileInfo           = "文件信息:"

	// 使用配置
	_logUseConfig = "使用配置:"
)

// 打印命令执行欢迎语
func WelcomeMessage(prefix string, message string) {
	log.Println(prefix+":", message)
}

// logUserConfigLocation 日志:输出配置文件位置
func (l Logs) ConfigLocation(message string) {
	log.Println(l.Prefix+":", _logConfigLocation, message)
}
func (l Logs) ConfigNotExist(message string) {
	log.Println(l.Prefix+":", _logConfigNotExist, message)
}
func (l Logs) ConfigIsExist(message string) {
	log.Println(l.Prefix+":", _logConfigIsExist, message)
}
func (l Logs) ConfigCreated(message string) {
	log.Println(l.Prefix+":", _logConfigCreated, message)
}
func (l Logs) ConfigCreatedFail(message string) {
	log.Println(l.Prefix+":", _logConfigCreatedFail, message)
}
func (l Logs) ConfigCreatedSuccess(message string) {
	log.Println(l.Prefix+":", _logConfigCreatedSuccess, message)
}
func (l Logs) ConfigCreatedDiscard(message string) {
	log.Println(l.Prefix+":", _logConfigCreatedDiscard, message)
}

// logUserConfigLocation 日志:输出用户配置文件位置
func (l Logs) UserConfigLocation(message string) {
	log.Println(l.Prefix+":", _logUserConfigLocation, message)
}
func (l Logs) UserConfigNotExist(message string) {
	log.Println(l.Prefix+":", _logUserConfigNotExist, message)
}
func (l Logs) UserConfigIsExist(message string) {
	log.Println(l.Prefix+":", _logUserConfigIsExist, message)
}
func (l Logs) UserConfigCreated(message string) {
	log.Println(l.Prefix+":", _logUserConfigCreated, message)
}
func (l Logs) UserConfigCreatedFail(message string) {
	log.Println(l.Prefix+":", _logUserConfigCreatedFail, message)
}
func (l Logs) UserConfigCreatedSuccess(message string) {
	log.Println(l.Prefix+":", _logUserConfigCreatedSuccess, message)
}
func UserConfigCreatedDiscard(prefix string, message string) {
	log.Println(prefix+":", _logUserConfigCreatedDiscard, message)
}

// logFileCreated 日志:文件已创建
func (l Logs) FileCreated(message string) {
	log.Println(l.Prefix+":", _logFileCreated, message)
}
func (l Logs) FileIsExist(message string) {
	log.Println(l.Prefix+":", _logFileIsExist, message)
}
func (l Logs) FileNotExits(message string) {
	log.Println(l.Prefix+":", _logFileNotExits, message)
}
func (l Logs) FileCreatedFail(message string) {
	log.Println(l.Prefix+":", _logFileCreatedFail, message)
}
func (l Logs) FileCreatedSuccess(message string) {
	log.Println(l.Prefix+":", _logFileCreatedSuccess, message)
}
func (l Logs) FileInfo(message string) {
	log.Println(l.Prefix+":", _logFileInfo, message)
}

// logUseConfig 日志:使用配置
func (l Logs) UseConfig(message string) {
	log.Println(l.Prefix+":", _logUseConfig, message)
}
