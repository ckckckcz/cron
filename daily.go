package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	commits := rand.Intn(4) + 5
	
	f, err := os.OpenFile("activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("gagal buka file:", err)
		os.Exit(1)
	}
	defer f.Close()

	for i := 1; i <= commits; i++ {
		now := time.Now().UTC().Format("2006-01-02T15:04:05")
		line := fmt.Sprintf("update %s - #%d\n", now, i)
		if _, err := f.WriteString(line); err != nil {
			fmt.Println("gagal nulis log:", err)
			os.Exit(1)
		}

		if err := run("git", "add", "activity.log"); err != nil {
			fmt.Println("gagal git add:", err)
			os.Exit(1)
		}

		date := time.Now().UTC().Format("2006-01-02")
		msg := fmt.Sprintf("chore: daily update #%d (%s)", i, date)
		if err := run("git", "commit", "-m", msg); err != nil {
			fmt.Println("gagal commit:", err)
			os.Exit(1)
		}
	}

	if err := run("git", "push"); err != nil {
		fmt.Println("gagal push:", err)
		os.Exit(1)
	}
}
