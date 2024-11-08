package tui

import (
	"os"
	"fmt"
	"time"
	"bufio"
	"strings"

	"github.com/ihor-vt/UsernameAvailabilityGo/pkg/style"
	"github.com/ihor-vt/UsernameAvailabilityGo/pkg/cursor"
	"github.com/ihor-vt/UsernameAvailabilityGo/pkg/social"
	"github.com/ihor-vt/UsernameAvailabilityGo/pkg/username"
	"github.com/ihor-vt/UsernameAvailabilityGo/pkg/components"
)

func Run() {
	fmt.Printf("%s\n%s\n", style.Cyan.Colorize(components.Header), components.Prompt)

	for {
		fmt.Printf("\n%s", components.TextInput)

		userInput := getUserInput()
		fmt.Print(cursor.Hide)
		fmt.Print(cursor.Up)
		fmt.Print(cursor.ClearLine)
		fmt.Print(cursor.Down)
		fmt.Printf(cursor.ClearLine)

		if userInput == "q" {
			fmt.Printf(cursor.ClearAfter)
			return
		}

		fmt.Printf("   Results for \"%s\" from %d platforms:\n\n", userInput, len(social.Platforms))

		t := time.Now()

		username.CheckAvailabilityConcurrent(userInput)

		fmt.Printf("\n    completed search in %v\n\n", time.Since(t))
		fmt.Print(strings.Repeat(cursor.Up, 25))
		fmt.Print(cursor.ClearLine)
		fmt.Print(cursor.Show)
	}
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
