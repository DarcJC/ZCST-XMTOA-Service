package main

import (
    "Onboarding/task"
    "context"
    "log"
    "time"
)

func main() {
    for {
        err := task.MainQueue.Consumer().Start(context.Background())
        if err != nil {
            log.Fatal(err)
        }
        time.Sleep(100 * time.Millisecond)
    }
}
