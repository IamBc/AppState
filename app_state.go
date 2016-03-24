package app_state

import  (
	"log"
	"time"
	)

func Hello(){
    log.Println(`Hello from AppState`)
}

type counter struct {
    value int
    updated_at Time
    updated_at_epoch int32 // int32(time.Now().Unix())
}


