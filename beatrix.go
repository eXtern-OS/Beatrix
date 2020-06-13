package beatrix

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	Token     string
	Issuer    string
	ChannelID string
	Discord   *discordgo.Session
	ErrorMode bool
)

// Function to fire up discord bot, issuer will be in the heading of message, token is bot token and channelId is channelID
func Init(issuer, token, channelID string) {
	Token = token
	Issuer = issuer
	ChannelID = channelID
	ErrorMode = false
	var err error
	Discord, err = discordgo.New("Bot " + Token)
	if err != nil {
		// Failed to init Beatrix
		log.Panic(err)
	}
	err = Discord.Open()
	if err != nil {
		// Failed to init Beatrix
		log.Panic(err)
	}
}

func Reinit() {
	var err error
	Discord, err = discordgo.New("Bot " + Token)
	if err != nil {
		// Failed to init Beatrix
		ErrorMode = false
		log.Println(err)
		return
	}
	err = Discord.Open()
	if err != nil {
		// Failed to init Beatrix
		ErrorMode = false
		log.Println(err)
		return
	}
	ErrorMode = false
	return
}

// Simply send a message to main channel
func Message(message string) {
	message = "[" + Issuer + "]\n" + message
	if ErrorMode {
		log.Println(message)
		Reinit()
		return
	}
	_, err := Discord.ChannelMessageSend(ChannelID, message)
	if err != nil {
		// Since we have goroutine, we don't have to return or something
		// Better re-init discord
		log.Println(err)
		Reinit()
	}
}

// Send error message to channel
func SendError(message, localIssuer string) {
	message = "[" + Issuer + " | " + localIssuer + "]\n" + message
	if ErrorMode {
		log.Println(message)
		Reinit()
		return
	}
	_, err := Discord.ChannelMessageSend(ChannelID, message)
	if err != nil {
		log.Println(err)
		Reinit()
	}
	return
}

func Panic(message string) {
	m := "[" + Issuer + " / PANIC]\n@everyone\n\n" + message
	if ErrorMode {
		log.Println(message)
		Reinit()
		return
	}
	_, err := Discord.ChannelMessageSend(ChannelID, m)
	if err != nil {
		log.Println(err)
		Reinit()
	}
	return
}
