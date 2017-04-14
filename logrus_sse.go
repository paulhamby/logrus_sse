package logrus_sse

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"gopkg.in/antage/eventsource.v1"
)

type SseHook struct {
	Writer eventsource.EventSource
}

//NewSseHook Creates a hook to be added to an instance of logger. This is called with
//hook, err := logrus_sse.NewSseHook(":8080")
//if err != nil {
//    log.Errorf("error adding hook %s", err)
//} else {
//    log.AddHook(hook)
//}
func NewSseHook(port string) (*SseHook, error) {
	es := eventsource.New(nil, nil)
	http.Handle("/log", es)
	go func() {
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
	return &SseHook{es}, nil
}

func (hook *SseHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	hook.Writer.SendEventMessage(line, "", "")
	return nil
}

func (hook *SseHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
