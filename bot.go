package twitchbot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gempir/go-twitch-irc"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	Username           string
	OauthPassword      string
	TwitchClientID     string
	TwitchClientSecret string
	tw                 *twitch.Client
	twitchChannels     []*TwitchChannel
}

func (b *Bot) AddChannel(twitchChannel *TwitchChannel) {
	b.twitchChannels = append(b.twitchChannels, twitchChannel)
}

func (b *Bot) Run() {
	log.Info("Starting twitch bot")

	// Create a twitch client
	b.tw = twitch.NewClient(b.Username, b.OauthPassword)

	// Graceful shutdown mechanism
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		<-gracefulStop
		log.Info("Shutting down")
		b.tw.Disconnect()
		os.Exit(0)
	}()

	// Add the channels we're listening to, and their respective triggers
	for _, twitchChannel := range b.twitchChannels {
		// Add the twitchChannel to the bot so we can listen.
		b.tw.Join((*twitchChannel).Name)

		// Add the twitchChannel's triggers
		b.addChanTriggers(twitchChannel)
	}

	// Connect to the twitch chat server
	err := b.tw.Connect()
	if err != nil {
		log.Panic(err)
	}
}

func (b *Bot) addChanTriggers(twitchChannel *TwitchChannel) {
	// Add the triggers for that twitchChannel
	//for _, trigger := range twitchChannel.Triggers {

	// b.tw.OnNewMessage(func(twitchChannel string, user twitch.User, message twitch.Message) {
	// 	fmt.Println(message.Text)
	// })
	//}

}
