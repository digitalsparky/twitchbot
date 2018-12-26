package main

import (
  "fmt"
  "os"

  "github.com/gempir/go-twitch-irc"
  _ "github.com/joho/godotenv/autoload"
  log "github.com/sirupsen/logrus"

)

func main() {
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
