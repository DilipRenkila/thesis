package main

import (
    "log"
    "golang.org/x/exp/inotify"
    "time"
)

func main() {
    watcher, err := inotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }

    done := make(chan bool)

    // Process events
    go func() {
        for {
            select {
            case ev := <-watcher.Event:
		    if ev.Mask&inotify.IN_MODIFY != inotify.IN_MODIFY {
			    continue
		    }
                log.Println("Sleep for 30 seconds ....")
                time.Sleep(time.Second*30)
            case err := <-watcher.Error:
                log.Println("error:", err)
            }
        }
    }()

    err = watcher.Watch("/mnt/LONTAS/ExpControl/dire15/status")
    if err != nil {
        log.Fatal(err)
    }

    // Hang so program doesn't exit
    <-done
    /* ... do stuff ... */
    watcher.Close()
}