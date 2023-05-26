package util

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/enescakir/emoji"
	"time"
)

func PrintLoading(suffix string, final string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[1], 100*time.Millisecond)
	s.Suffix = " " + suffix + "\n"
	s.FinalMSG = fmt.Sprintf("%v %v\n", emoji.CheckMark, final)
	s.Start()
	return s
}

func PrintErrAndReturn(s string) error {
	fmt.Printf("%v %v\n", emoji.ExclamationMark, s)
	return nil
}
