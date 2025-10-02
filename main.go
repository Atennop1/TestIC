package main

import (
	"context"
	"encoding/json"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
	"math/rand"
	"os"
	"os/signal"
)

var linesObject struct {
	Lines []string `json:"lines"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(DefaultHandler),
	}

	b, err := bot.New(os.Getenv("BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}

	linesJson, err := os.ReadFile("lines.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(linesJson, &linesObject)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "start", bot.MatchTypeCommand, StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "swaga", bot.MatchTypeCommand, SwagaHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "author", bot.MatchTypeCommand, AuthorHandler)

	b.Start(ctx)
}

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	sendMessage(ctx, b, update.Message.Chat.ID, "Unknown command. I know only /start, /swaga and /author!")
}

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	sendMessage(ctx, b, update.Message.Chat.ID, "Welcome to this bot! Type /swaga for random line from ICEGERGERT songs or /author to get info about my creator")
}

func SwagaHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	sendMessage(ctx, b, update.Message.Chat.ID, "Random line from ICEGERGERT songs:\n"+linesObject.Lines[rand.Intn(len(linesObject.Lines))])
}

func AuthorHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	sendMessage(ctx, b, update.Message.Chat.ID, "Author:\ngroup БИ25-6\n@atennop\ngithub.com/Atennop1")
}

func sendMessage(ctx context.Context, b *bot.Bot, chatId int64, message string) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatId,
		Text:   message,
	})

	if err != nil {
		panic(err)
	}
}
