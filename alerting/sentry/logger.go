package sentry

import (
	"fmt"

	apex "github.com/francoishill/log"
	"github.com/francoishill/log/handlers/level"
	raven "github.com/getsentry/raven-go"
)

//NewLogHandler returns a new logger that logs to sentry
func NewLogHandler(client *raven.Client, logLevel apex.Level) apex.Handler {
	sentryLogger := &logger{client: client}
	return level.New(sentryLogger, logLevel)
}

type logger struct {
	client *raven.Client
}

var severityMapping = [...]raven.Severity{
	apex.DebugLevel:     raven.DEBUG,
	apex.InfoLevel:      raven.INFO,
	apex.NoticeLevel:    raven.WARNING,
	apex.WarnLevel:      raven.WARNING,
	apex.ErrorLevel:     raven.ERROR,
	apex.CriticalLevel:  raven.FATAL,
	apex.AlertLevel:     raven.FATAL,
	apex.EmergencyLevel: raven.FATAL,
}

func (l *logger) HandleLog(e *apex.Entry) error {
	packet := raven.NewPacket(e.Message)
	packet.Init("") //will setup default fields like project, culprit, etc
	packet.Level = severityMapping[e.Level]

	tags := map[string]string{}
	for key, val := range e.Fields {
		tags[key] = fmt.Sprintf("%v", val)
		packet.Extra["full-"+key] = fmt.Sprintf("%+v", val)
	}

	go l.client.Capture(packet, tags)

	return nil
}
