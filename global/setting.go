package global

import (
	"github.com/dagou8/blog-service/pkg/logger"
	"github.com/dagou8/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
