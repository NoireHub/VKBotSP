package handlers

const msgHelp = `Я бот

Могу сохранять твои цитаты и присылать случайные из них.
Также у меня есть кнопки обо мне и о моем создателе`

const msgHello = "Привет! 👾\n\n" + msgHelp

const (
	msgUnknownCommand = "Неизвестная комманда 🤔"
	msgError   = "Произошла ошибка 🙊"
	msgNoSavedQuotes   = "У вас еще нет сохраненных цитат 🙊"
	msgSaved          = "Успешно сохранено! 👌"
	msgAlreadyExists  = "Такая цитата уже сохранена 🤗"
	msgReadyToSave = "Готов сохранить цитату"
	msgNoLucky = "Повезет в следующий раз"
	msgLucky = "Тебе везет"
	msgSuperLucky = "Такое везение сравнимо только с божетсвами"
)

var Prophesy = []string{
	"Предсказание 1",
	"Предсказание 2",
}