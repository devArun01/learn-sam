package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/devArun01/learn-sam/pkg"
	"github.com/manifoldco/promptui"
)

func test() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n %s", result, pkg.Name())
}

func main() {
	lambda.Start(test)
	// test()
}
