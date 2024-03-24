package utils

import (
	"gorman/m3u8dl/models"
	"testing"
)

func TestDownload(t *testing.T) {
	downloadDir := "D:\\video_download"
	param := models.DL{
		Fan: "来玩游戏吧",
		Episodes: []models.Episode{
			{
				Url:   "https://vip.ffzy-online1.com/20221214/6508_c6ccfa67/index.m3u8?=1711283808000",
				Type:  "m3u8",
				Title: "第一集",
			},
			// {
			// 	Url:   "https://vip.ffzy-online1.com/20221214/6509_fd52add8/index.m3u8?=1711283809000",
			// 	Type:  "m3u8",
			// 	Title: "第二集",
			// },
		},
	}
	if err := Download(param, downloadDir); err != nil {
		t.Error(err)
	}
}
