package bot

import (
  "fmt"
  "os"

  "github.com/gempir/go-twitch-irc"
  _ "github.com/joho/godotenv/autoload"
  log "github.com/sirupsen/logrus"

  "github.com/digitalsparky/twitchbot/lib/channel"
)

type Bot struct {
  channels []Channel

}

func New() {
  client := twitch.NewClient(os.Getenv("BOT_USERNAME"), os.Getenv("OAUTH_PASSWORD"))

  client.OnNewMessage(func(channel string, user twitch.User, message twitch.Message) {
    fmt.Println(message.Text)
  })

  client.Join("digitalsparky")

  err := client.Connect()
  if err != nil {
    log.Panic(err)
  }
}
