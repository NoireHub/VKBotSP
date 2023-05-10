package main

import (
	"os"

	"github.com/NoireHub/VKBotSP/internal/botVK"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/sirupsen/logrus"
)

func main() {
	token:= os.Getenv("token")
	dataBaseURL:= os.Getenv("dbURL")

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

//1 кнопка - о создателе => 2 кнопки github leetcode