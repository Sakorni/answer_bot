package main

import (
	"CAOS/utils"
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const adminId = 277858809
const ErrorResponse = "Что-то пошло не так. Попробуйте ещё раз (советую изменить входные данные)"
const NoContentResponse = "По запросу ничего не найдено. Попробуйте написать как-нибудь иначе..."
const ToMuchContetnF = `Бля, чёт дохуя выдаёт, уточни запрос плз
(Вы только что были спасены от града сообщений в размере %d штук)`
const NoFuckYou = "Это что за попытки гриферства?\nСкоро в бан полетишь, дружок-пирожок (Как только прикручу оный)"
const GrifferF = "Это пидор гриферит!\nОн отправлял следующие непристойности:\n%s"

type ReportContent struct {
	User    tgbotapi.User
	Caption string
}

func (rc ReportContent) UserDescription() string {
	return fmt.Sprintf("{%s}:%s, {%s}:%s, {%s}:%s",
		"Username", rc.User.UserName,
		"First name", rc.User.FirstName,
		"Last name", rc.User.LastName,
	)
}

func main() {
	key := os.Getenv("BOT_KEY")
	if key == "" {
		log.Fatal("Fatal error! No BOT_KEY variable provided!")
	}
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		text := strings.TrimSpace(update.Message.Text)
		if deprecated(text) {
			report := ReportContent{
				*update.Message.From,
				fmt.Sprintf(GrifferF, text),
			}
			bot.Send(reportToAdmin(report))

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, NoFuckYou)
			bot.Send(msg)
			continue
		}
		var response []string
		answer, err := utils.AskForAnswer(text)

		switch err {
		case nil:
			response = answer
		case utils.NoResult:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, NoContentResponse)
			bot.Send(msg)
		default:
			report := ReportContent{
				*update.Message.From,
				err.Error(),
			}
			bot.Send(reportToAdmin(report))
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, ErrorResponse)
			bot.Send(msg)
		}

		if len(response) > 10 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				fmt.Sprintf(ToMuchContetnF, len(response)))
			bot.Send(msg)
			continue
		}
		for _, r := range response {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, r)
			bot.Send(msg)
		}
	}
}

func deprecated(s string) bool {
	return len(s) < 3 ||
		strings.Contains(s, "_") ||
		strings.Count(s, "%") > 1
}

func reportToAdmin(content ReportContent) tgbotapi.MessageConfig {
	message := fmt.Sprintf(
		"Произошла ошибка!\nИсточник - {%s}\nСодержание ошибки - {%s}",
		content.UserDescription(), content.Caption)
	return tgbotapi.NewMessage(adminId, message)
}
