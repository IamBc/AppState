package app_state

import  (
	"log"
	"time"
	)

func Hello(){
    log.Println(`Hello from AppState`, int32(time.Now().Unix()))
}

type counter struct {
    value int
    updated_at time.Time
    updated_at_epoch int32 // int32(time.Now().Unix())
}


type flag struct {
    is_up bool
    updated_at time.Time
    updated_at_epoch int32 // int32(time.Now().Unix())
}

type status struct {
    status int
    statusLabel string
    updated_at time.Time
    updated_at_epoch int32 // int32(time.Now().Unix())
    statusLabels map[int] string
}

type appState struct {
    Counters map[string] counter
    Flags map[string] flag
    Statuses map[string] status
    StatusFilePath string
}

