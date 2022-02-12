package main

import (
	"log"
	"os/exec"
	"time"

	"github.com/Netflix/go-expect"
)

func main() {
	c, err := expect.NewConsole()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	cmd := exec.Command("gh", "pr", "create", "--fill", "--base", "happy-path")
	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()

	go func() {
		c.ExpectEOF()
	}()

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)
	c.Send("iHello world\x1b")
	c.SendLine(":wq")

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
