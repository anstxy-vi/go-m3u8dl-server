package v1

import (
	"fmt"
	"gorman/m3u8dl/models"
	"gorman/m3u8dl/pkg/utils"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// 命令路径
const ClI_PATH = "D:\\Program Files\\N_m3u8DL-CLI_v3.0.2.exe"

// 存放目录
const SAVE_DIR = "Downloads"

type DLService struct{}

func NewDLApi() *DLService {
	return &DLService{}
}

func descDir(param models.DL) string {
	return filepath.Join(SAVE_DIR, param.Fan)
}

func (dl *DLService) HandleUrls(ctx *gin.Context) {
	params := models.DL{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	if err := dl.checkM3u8Cli(); err != nil {
		ctx.AbortWithStatus(http.StatusNotImplemented)
	}
	if err := dl.initDir(params); err != nil {
		ctx.AbortWithStatus(http.StatusNotImplemented)
	}
	if err := dl.download(params); err != nil {
		ctx.AbortWithStatus(http.StatusNotImplemented)
	}
}

func (dl *DLService) checkM3u8Cli() error {
	fmt.Println("Check m3u8 cli...")

	command := fmt.Sprintf("%s", ClI_PATH)
	cmd := exec.Command(command)
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("Check m3u8 cli completed")
	return nil
}

// 初始化目录
func (dl *DLService) initDir(params models.DL) error {
	err := os.RemoveAll(descDir(params))
	if err != nil {
		return err
	}
	utils.MkdirWhenNotExist(descDir(params))
	return nil
}

// 开始循环下载
func (dl *DLService) download(param models.DL) error {
	fmt.Printf("总共%v个视频\n", len(param.Episodes))
	for index, episode := range param.Episodes {
		seq := index + 1
		if len(param.Episodes) != seq {
			fmt.Printf("下载中...(当前第%v个视频, 还有%v个视频等待下载)...\n", seq, len(param.Episodes)-seq)
		} else {
			fmt.Println("下载中...(即将完成)")
		}

		var commands = []string{
			ClI_PATH,
			episode.Url,
			"--workDir",
			descDir(param),
			"--saveName",
			episode.Title,
		}
		cmd := exec.Command(commands[0], commands[1:]...)

		if err := cmd.Run(); err != nil {
			fmt.Println("请检查网络是否能连接m3u8地址")
			return err
		}

		if o, err := cmd.Output(); err == nil {
			fmt.Println(string(o))
		}
		os.RemoveAll(filepath.Join(descDir(param), episode.Title))
		fmt.Printf("第%v个视频下载完成\n", seq)
	}
	fmt.Printf("[%s]已全部下载完成\n", param.Fan)
	return nil
}
