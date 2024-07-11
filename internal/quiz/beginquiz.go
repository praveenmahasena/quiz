package quiz

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/praveenmahasena/quiz/internal/file"
)

func Start(csv *file.CSV) error {
	s := bufio.NewScanner(csv.File)
	var score int

	s.Split(bufio.ScanLines)
	for s.Scan() {
		if errors.Is(s.Err(), io.EOF) {
			fmt.Printf("your score is %v", score)
			return nil
		}

		question, ans, err := format(s.Text())

		var a string

		if err != nil {
			return err
		}

		fmt.Print(question + "=")

		fmt.Fscan(os.Stdin, &a)

		if ans != a {
			fmt.Printf("possible wrong answer %v", score)
			return nil
		}

		score += 1
	}

	return nil
}

func format(str string) (string, string, error) {
	leng := len(str)
	question := ""
	ans := ""

	for i := 0; i < leng; i++ {
		if str[i] != ',' {
			question += string(str[i])
			continue
		}
		ans = str[i+1:]
		break
	}

	if question == ans {
		return "", "", fmt.Errorf("error during parsing input might be a format error in CSV file")
	}
	return question, ans, nil
}
