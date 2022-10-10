package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

var lock sync.Mutex

func execCmd(cmd []string, msg Message) error {
	lock.Lock()
	defer lock.Unlock()
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	env := os.Environ()
	env = append(env, fmt.Sprintf("PUSHOVER_MESSAGE=%s", msg.Message))
	env = append(env, fmt.Sprintf("PUSHOVER_TITLE=%s", msg.Title))
	c.Env = env
	return c.Run()
}
