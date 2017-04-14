### Logrus_sse

This is a hook for logrus which will start up a web server on a configurable port and will stream your logs via EventSource/Server-Sent Events.

It's very experimental at the moment, but works for my use case.

### Example Usage

```
func init() {
	hook, err := logrus_sse.NewSseHook(":8080")
	if err != nil {
		log.Errorf("error adding hook %s", err)
	} else {
		log.AddHook(hook)
	}
	log.SetFormatter(&log.TextFormatter{DisableColors: true})
}

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		log.WithField("LogEntry", i).Info("Something interesting..")
	}
}
```

You can then connect to your host on port 8080 and you should see some events.