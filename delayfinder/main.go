package main

import (
    "log"
    "golang.org/x/exp/inotify"
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
                log.Println("Executing ....")
                capshow()
            case err := <-watcher.Error:
                log.Println("error:", err)
            }
        }
    }()

    err = watcher.Watch("/mnt/LONTAS/ExpControl/dire15/info")
    if err != nil {
        log.Fatal(err)
    }

    // Hang so program doesn't exit
    <-done
    /* ... do stuff ... */
    watcher.Close()
}