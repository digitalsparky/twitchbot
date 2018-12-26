package twitchbot

import (
  "fmt"
  "os"

  "github.com/gempir/go-twitch-irc"
  log "github.com/sirupsen/logrus"

)

type Bot struct {
  Username	string
  OauthPassword string
  TwitchClientID string
  TwitchClientSecret string
  tc twitch.Client
  //channels []Channel

}

func (b *Bot) Start() {
  b.tc = twitch.NewClient()



  b.tc.OnNewMessage(func(channel string, user twitch.User, message twitch.Message) {
    fmt.Println(message.Text)
  })

  b.tc.Join("digitalsparky")

  err := b.tc.Connect()
  if err != nil {
    log.Panic(err)
  }
}
