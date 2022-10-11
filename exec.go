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
	switch msg.Priority {
	case -2:
		env = append(env, "PUSHOVER_PRIORITY=-2")
		env = append(env, "PUSHOVER_PRIORITY_STR=lowest")
	case -1:
		env = append(env, "PUSHOVER_PRIORITY=-1")
		env = append(env, "PUSHOVER_PRIORITY_STR=low")
	case 0:
		env = append(env, "PUSHOVER_PRIORITY=0")
		env = append(env, "PUSHOVER_PRIORITY_STR=normal")
	case 1:
		env = append(env, "PUSHOVER_PRIORITY=1")
		env = append(env, "PUSHOVER_PRIORITY_STR=high")
	case 2:
		env = append(env, "PUSHOVER_PRIORITY=2")
		env = append(env, "PUSHOVER_PRIORITY_STR=emergency")
	}
	c.Env = env
	return c.Run()
}
