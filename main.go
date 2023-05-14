package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5253855724966-5258361487205-YawGfJnAQGSacuh7Y1Zpbnzi")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A057GL291LN-5273929263185-0ae0de9d835e9391f38abcad559a531a850f3947b62eaa9ba46eb39c72985fd2")

	bot := slacker.NewClient(os.Getenv("SLACKER_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition) {
		Description: "yob calculator", 
		Example: "my yob is 2023",
		Handler: func (botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter)  {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2022-yob
			r :=  fmt.Sprintf("age is %d", age)
			response.Reply(r)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
