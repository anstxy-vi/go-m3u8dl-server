package initial

import (
	"path/filepath"

	"gorman/m3u8dl/config"
	"gorman/m3u8dl/pkg/logger"
	"gorman/m3u8dl/pkg/utils"
)

// 初始化日志
func InitLogger() {
	cwd, err := utils.CWD()
	if err != nil {
		panic(err)
	}
	conf := config.Global
	err = logger.InitDefault(
		filepath.Join(cwd, conf.Logger.Dir, conf.Logger.Access),
		filepath.Join(cwd, conf.Logger.Dir, conf.Logger.Error),
		filepath.Join(cwd, conf.Logger.Dir, conf.Logger.Debug),
		conf.Debug,
	)

	if err != nil {
		panic(err)
	}

	println("初始化 Logger 完成")

}
