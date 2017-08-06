package cakewalk

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCakewalk(t *testing.T) {
	Cakewalk()
	subproc := exec.Command(os.Args[0])
	stdin, _ := subproc.StdinPipe()
	stdout, _ := subproc.StdoutPipe()
	defer stdin.Close()

	input1 := "3\n"
	input2 := "1 3 2\n"

	subproc.Start()
	io.WriteString(stdin, input1)
	io.WriteString(stdin, input2)
	reader := bufio.NewReader(stdout)
	bytes, _, _ := reader.ReadLine()
	output := string(bytes)
	assert.Equal(t, "11", output, "not equqal")

}
