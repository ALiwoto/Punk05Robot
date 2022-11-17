package tests

import (
	"testing"

	"github.com/AnimeKaizoku/ssg/ssg"
)

func TestSplitCommand(t *testing.T) {
	const myCommand = "/setFooter -100123456 ANY TEXT HERE \nAND HERE"
	myStrs := ssg.SplitN(myCommand, 3, " ")
	print(myStrs)

	const myCommand2 = "/setFooter 12345"
	myStrs = ssg.SplitN(myCommand2, 3, " ")
	print(myStrs)
}
