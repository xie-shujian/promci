package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

var Log1 *log.Logger
var Conf *Config

func InitLogger() {
	var folder string
	if runtime.GOOS == "windows" {
		folder, _ = os.Getwd()
	} else {
		folder = "/var/log"
	}
	fileName := folder + "/promci.log"

	logWriter, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic("create log file failed" + err.Error())
	}
	Log1 = log.New(logWriter, "", log.LstdFlags)
}

func ReadConfig() {
	var folder string
	if runtime.GOOS == "windows" {
		folder, _ = os.Getwd()
	} else {
		folder = "/etc/promci"
	}
	fileName := folder + "/promci.yml"

	Conf = &Config{}
	dataBytes, err := os.ReadFile(fileName)
	if err != nil {
		Log1.Panicln("read file failed", err)
	}
	err = yaml.Unmarshal(dataBytes, Conf)
	if err != nil {
		Log1.Panicln("parser yaml failed", err)
	}
}

func BuildGitClone(url1 string, dir1 string) string {
	s := "git clone " + url1 + " " + dir1
	ms := Mask(s)
	Log1.Println(ms)
	return s
}
func BuildGitPull(url1 string, dir1 string) string {
	s := "git --work-tree=" + dir1 + " --git-dir=" + dir1 + "/.git pull " + url1
	ms := Mask(s)
	Log1.Println(ms)
	return s
}

func RunGitCmd(git_cmd string) {
	var cmd1 *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd1 = exec.Command("cmd", "/C", git_cmd)
	} else {
		cmd1 = exec.Command("sh", "-c", git_cmd)
	}
	output, err := cmd1.Output()
	if err != nil {
		Log1.Println(err)
	} else {
		Log1.Println(string(output))
	}
}

func Mask(s string) string {
	i1 := strings.Index(s, "oauth2:")
	i2 := strings.Index(s, "@")
	if i1 > 0 && i2 > 0 {
		return s[:i1] + "******" + s[i2:]
	}
	return s
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Gitlab-Token")
	if Conf.PromciAccessToken != token {
		Log1.Println("token not match")
		return
	}

	repositoryName := r.URL.Query().Get("repository")
	if !Conf.ExistRepository(repositoryName) {
		Log1.Println("Repository not found")
		return
	}
	url1, dir1 := Conf.BuildRepositoryAccessUrl(repositoryName)

	git_pull := BuildGitPull(url1, dir1)
	go RunGitCmd(git_pull)
}
