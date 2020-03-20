package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var n = 10
var quickestTime time.Duration

func main() {
	for {
		guess()
	}
}

func guess() {
	start := time.Now()
	fmt.Println("Guessing Game:")
	fmt.Printf("- guess a number from 1 to %d\n", n)
	say(fmt.Sprintf("guess from 1 to %d", n))

	rand.Seed(time.Now().Unix())
	number := rand.Intn(n)
	number++

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
			say("hello yourself!")
			continue
		}

		userNumber, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("🤮 that's not a number!")
			say("what? number please")
			continue
		}

		if userNumber == number {
			fmt.Println("congrats! you guessed the number 🎆😀")
			say("you guessed it!")

			duration := time.Since(start)
			fmt.Printf("it took you %s\n", duration)
			if quickestTime == 0 {
				quickestTime = duration
			}
			if quickestTime > duration {
				quickestTime = duration
				fmt.Println("that's your fastest game yet")
			}

			text, _ = reader.ReadString('\n')

			fmt.Println("\n\n\n-------------------------")
			fmt.Println("New Game!")
			fmt.Println("-------------------------")
			return
		}

		if userNumber > number {
			fmt.Println("🔻 that's too high. guess a smaller number")
			say("too high")
			continue
		}
		if userNumber < number {
			fmt.Println("🔺 that's too low. guess a bigger number")
			say("too low")
			continue
		}
	}
}

// say uses the say command to speak
// TODO: use an actual golang library to do this instead of depending on a command
func say(s string) {
	cmd := exec.Command("say", s)
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
