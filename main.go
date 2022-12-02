package main


import(
	"fmt"
	"github.com/UgoRastell/VALOGO/bot"
	"github.com/UgoRastell/VALOGO/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil{
		fmt.Printf(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}