package bot

import(
	"fmt"
	"github.com/UgoRastell/VALOGO/config"
	"github.com/bwmarrin/discordgo"
)

var botId string
var goBot *discordgo.Session


func Start()  {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil{
		fmt.Println(err.Error())
	}

	botId = u.ID

	goBot.AddHandler(messageHandler)
	goBot.AddHandler(apiGet)
	
	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Bot is running...")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate){
		
	if m.Author.ID == botId{
		return
	}

	if m.Content == "pipi"{
		_, _=s.ChannelMessageSend(m.ChannelID, "caca")
	}
}

var client *http.Client

func apiGet()  {
	client := &http.Client{Timeout: 10 * time.Second}
}


