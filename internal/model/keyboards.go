package model

import "github.com/SevereCloud/vksdk/v2/object"

func NewMainMenuKeyboard() object.MessagesKeyboard {
	KeyBoard := object.MessagesKeyboard{}

	KeyBoard.AddRow()
	KeyBoard.AddTextButton("Рандом", "", "primary")
	KeyBoard.AddTextButton("Цитаты", "", "primary")

	KeyBoard.AddRow()
	KeyBoard.AddTextButton("О боте", "", "primary")
	KeyBoard.AddTextButton("Об авторе", "", "primary")

	return KeyBoard
}

func NewAboutAuthorKeyboard() object.MessagesKeyboard {
	KeyBoard := object.MessagesKeyboard{}

	KeyBoard.AddRow()
	KeyBoard.AddOpenLinkButton("https://github.com/NoireHub", "Github", "")
	KeyBoard.AddOpenLinkButton("https://leetcode.com/NoireHub/", "Lettcode", "")

	KeyBoard.AddRow()
	KeyBoard.AddTextButton("Назад", "", "negative")

	return KeyBoard
}

func NewAboutBotKeyboard() object.MessagesKeyboard {
	KeyBoard := object.MessagesKeyboard{}

	KeyBoard.AddRow()
	KeyBoard.AddTextButton("Возможности", "", "primary")

	KeyBoard.AddRow()
	KeyBoard.AddTextButton("Назад", "", "negative")

	return KeyBoard
}

func NewQuoteKeyboard() object.MessagesKeyboard {

	KeyBoard := object.MessagesKeyboard{}

	KeyBoard.AddRow()
	KeyBoard.AddTextButton("Добавить цитату", "", "primary")
	KeyBoard.AddTextButton("Получить случайную цитату", "", "primary")

	KeyBoard.AddRow()
	KeyBoard.AddTextButton("Назад", "", "negative")

	return KeyBoard
}

func NewRandomMenuKeyboard() object.MessagesKeyboard {

	KeyBoard := object.MessagesKeyboard{}

	KeyBoard.AddRow()
	KeyBoard.AddTextButton("Покрутить колесо удачи", "", "primary")
	KeyBoard.AddTextButton("Получить предсказание", "", "primary")

	KeyBoard.AddRow()
	KeyBoard.AddTextButton("Назад", "", "negative")

	return KeyBoard
}
