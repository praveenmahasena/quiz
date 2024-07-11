package internal

import (
	"flag"

	"github.com/praveenmahasena/quiz/internal/file"
	"github.com/praveenmahasena/quiz/internal/quiz"
)

func Run() error {
	path := flag.String("path", "", "path to your coustomized questions.csv")

	flag.Parse()

	cvs, cvsErr := file.LoadCSV(*path)
	if cvsErr != nil {
		return cvsErr
	}

	return quiz.Start(cvs)
}
