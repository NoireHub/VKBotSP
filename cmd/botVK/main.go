package main

import (
	"os"
	"flag"

	"github.com/NoireHub/VKBotSP/internal/botVK"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/sirupsen/logrus"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath,"config-path","configs/bot.toml","path to config file")

}


func main() {
	var token string
	var dataBaseURL string

	flag.Parse()
	config:= botvk.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		token = os.Getenv("TOKEN")
		dataBaseURL = os.Getenv("DB_URL")
	}else{
		token = config.Token
		dataBaseURL = config.DatabaseURL + os.Getenv("HOST")
	}

	logrus.Info(os.Getenv("HOST"))

	vk := api.NewVK(token)
	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		logrus.Fatal("Cannot get group ID: " + err.Error())
	}
	
	if err:= botvk.Start(dataBaseURL, vk, group[0].ID);
	err != nil {
		logrus.Fatal(err.Error())
	}
}