package menu

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

func selectFromMenu(prompt promptui.Select) (int, string) {

	key, value, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(123)
	}

	return key, value
}
