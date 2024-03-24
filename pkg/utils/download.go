package utils

import (
	"fmt"
	"gorman/m3u8dl/models"
	"os"
	"os/exec"
	"path/filepath"
)

const CLI_PATH = "D:\\Program Files\\N_m3u8DL-CLI_v3.0.2.exe"

func Download(param models.DL, downloadDir string) error {
	fmt.Printf("总共%v个视频\n", len(param.Episodes))

	descDir := filepath.Join(downloadDir, param.Fan)
	MkdirWhenNotExist(descDir)
	for index, episode := range param.Episodes {
		seq := index + 1
		if len(param.Episodes) != seq {
			fmt.Printf("下载中...(当前第%v个视频, 还有%v个视频等待下载)...\n", seq, len(param.Episodes)-seq)
		} else {
			fmt.Println("下载中...(即将完成)")
		}

		var commands = []string{
			CLI_PATH,
			episode.Url,
			"--workDir",
			descDir,
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
		os.RemoveAll(filepath.Join(descDir, episode.Title))
		fmt.Printf("第%v个视频下载完成\n", seq)
	}
	fmt.Printf("[%s]已全部下载完成\n", param.Fan)
	return nil
}
