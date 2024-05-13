package services

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"helper_bot/internal/client"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	StartCmd   = "/start"
	HelpCmd    = "/help"
	LearnCmd   = "/learn"
	SettingCmd = "/setting"
)

type RunBot interface {
	RunBot(token string) error
}

type StartBotService struct {
}

func (s *StartBotService) RunBot(token, openWeatherMapAPIKey string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		reply := ""
		switch update.Message.Text {
		case StartCmd:
			reply = client.MsgHello
		case HelpCmd:
			reply = client.MsgHelp
		case LearnCmd:
			reply = client.MsgLearning
		case SettingCmd:
			reply = client.MsgSetting
		default:
			weather, err := getWeather(update.Message.Text, openWeatherMapAPIKey)
			if err != nil {
				reply = "Город не найден."
			} else {
				reply = weather
			}
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	}

	return nil
}

func getWeather(city, openWeatherMapAPIKey string) (string, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, openWeatherMapAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data client.WeatherData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	fmt.Println(data)

	return fmt.Sprintf("Город: %s\nТемпература: %.2fC\nОщущается: %.2fC\nСкорость ветра: %.2fм/с\nВлажность: %.2f%%\nНебо: %s\nhttp://openweathermap.org/img/w/%s.png",
		data.Name, data.Main.Temp, data.Main.FeelsLike, data.Wind.Speed, data.Main.Humidity, data.Weather[0].Main, data.Weather[0].Icon), nil
}
