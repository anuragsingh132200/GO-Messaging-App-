package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type User struct {
	ID   int
	Name string
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Charlie"},
}

type Message struct {
	SenderID   int
	ReceiverID int
	Content    string
	Timestamp  time.Time
}

var messageLogs []Message

func main() {
	fmt.Println("Welcome to Messaging App")

	for {
		fmt.Println("1. Send message between two users")
		fmt.Println("2. Broadcast message to all users")
		fmt.Println("3. View message logs of a user")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			sendMessage()
		case 2:
			broadcastMessage()
		case 3:
			viewMessageLogs()
		case 4:
			fmt.Println("Exiting...")
			displayAllLogs()
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func sendMessage() {
	var senderID, receiverID int
	var content string

	fmt.Print("Enter sender ID: ")
	fmt.Scanln(&senderID)
	fmt.Print("Enter receiver ID: ")
	fmt.Scanln(&receiverID)
	fmt.Print("Enter message content: ")
	content = readMultiWordInput()

	if content = strings.TrimSpace(content); content == "" {
		content = getRandomFact()
	}

	newMessage := Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		Timestamp:  time.Now(),
	}

	messageLogs = append(messageLogs, newMessage)
	fmt.Printf("User %d sent message to User %d: %s\n", senderID, receiverID, content)
}

func broadcastMessage() {
	var senderID int
	var content string

	fmt.Print("Enter sender ID: ")
	fmt.Scanln(&senderID)
	fmt.Print("Enter message content: ")
	content = readMultiWordInput()

	if content = strings.TrimSpace(content); content == "" {
		content = getRandomFact()
	}

	for _, u := range users {
		sendMessageHelper(senderID, u.ID, content)
	}
}

func sendMessageHelper(senderID, receiverID int, content string) {
	newMessage := Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		Timestamp:  time.Now(),
	}

	messageLogs = append(messageLogs, newMessage)
	fmt.Printf("User %d sent message to User %d: %s\n", senderID, receiverID, content)
}

func viewMessageLogs() {
	var userID int

	fmt.Print("Enter user ID to view logs: ")
	fmt.Scanln(&userID)

	for _, msg := range messageLogs {
		if msg.SenderID == userID || msg.ReceiverID == userID {
			fmt.Printf("User %d: %s\n", msg.SenderID, msg.Content)
		}
	}
}

func displayAllLogs() {
	fmt.Println("Message Logs:")
	for _, msg := range messageLogs {
		fmt.Printf("Sender: %d, Receiver: %d, Content: %s\n", msg.SenderID, msg.ReceiverID, msg.Content)
	}
}

type Quote struct {
	Quote string `json:"quote"`
}

func getRandomFact() string {
	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		return "Error fetching quote"
	}
	defer resp.Body.Close()

	var quote Quote
	err = json.NewDecoder(resp.Body).Decode(&quote)
	if err != nil {
		return "Error parsing quote"
	}

	return quote.Quote
}

func readMultiWordInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
