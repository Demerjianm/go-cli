package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"gopkg.in/urfave/cli.v2"
)

func sendText(number string, message string) {
	fmt.Println("Sending text logic below here.")
	fullNumber := fmt.Sprint("+1", number)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if number == "" {
		fmt.Println("No Phone number!")
		return
	}
	if message == "" {
		fmt.Println("No message to send.")
		return
	}
	if err != nil {
		fmt.Println("NewSession error:", err)
		return
	}

	client := sns.New(sess)
	input := &sns.PublishInput{
		Message:     aws.String("Hello world!"),
		PhoneNumber: aws.String(fullNumber),
	}

	result, err := client.Publish(input)
	if err != nil {
		fmt.Println("Publish error:", err)
		return
	}

	fmt.Println(result)
}

func handle() {
	fmt.Println("What number would you like to text?")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	fmt.Println("What do you want to say?")
	var message string
	fmt.Scanln(&message)
	fmt.Println(message)

	sendText(input, message)
}

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the lineliness!",
		Action: func(c *cli.Context) error {
			handle()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
