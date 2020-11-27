package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
		Message:     aws.String(message),
		PhoneNumber: aws.String(fullNumber),
	}

	result, err := client.Publish(input)
	if err != nil {
		fmt.Println("Publish error:", err)
		return
	}

	fmt.Println(result)
}

// ask user single line question
func readSingleLine(question string) string {
	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return "Error"
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	fmt.Println(input)
	return input
}

func handle() {
	fmt.Println("What number would you like to text?")
	var input string
	fmt.Scanln(&input)
	message := readSingleLine("What do you want to say?")
	sendText(input, message)
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "text-message",
				Usage: "fight the loneliness!",
				Action: func(c *cli.Context) error {
					handle()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
