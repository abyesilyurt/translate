package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var request TranslationRequest
	if len(os.Args) == 2 {
		request.Filename = os.Args[1]
		request.TargetLang = "en"
	} else if len(os.Args) == 3 {
		request.Filename = os.Args[2]
		request.TargetLang = os.Args[1]
	} else {
		fmt.Println("Usage: translate <lang> <filename>")
		return
	}
	bytes, err := TranslateDocument(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputFilename := fmt.Sprintf("%s.translated.%s.pdf", request.Filename, request.TargetLang)
	err = os.WriteFile(outputFilename, bytes, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputFilename)
	_ = exec.Command("open", outputFilename).Run()
}
