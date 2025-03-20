package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerURL  string `yaml:"domain"`
	APIToken   string `yaml:"token"`
	Visibility string `yaml:"visibility"`
}

func main() {
	// 設定ファイルを読み込む
	config, err := loadConfig("ihhi.yaml")
	if err != nil {
		fmt.Println("設定ファイル読み込みエラー:", err)
		return
	}

	// 投稿する内容を定義
	postData := map[string]string{
		"visibility": config.Visibility,
		"i":          config.APIToken,
		"text":       ":neko_neru_nya:ｲｯﾋ",
	}

	// JSONに変換
	jsonData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println("JSON変換エラー:", err)
		return
	}

	// HTTPリクエストを作成
	resp, err := http.Post("https://"+config.ServerURL+"/api/notes/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("リクエストエラー:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("投稿成功！")
	} else {
		fmt.Println("投稿失敗。ステータスコード:", resp.StatusCode)
	}
}

func loadConfig(filename string) (Config, error) {
	var config Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return config, fmt.Errorf("ファイル読み込みエラー: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("YAMLパースエラー: %v", err)
	}

	return config, nil
}
