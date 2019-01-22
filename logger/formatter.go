package logger

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type TemplateFormatter struct{}

func (fmter *TemplateFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timeLayout := time.RFC1123
	if len(entry.Data) > 0 {
		dataBuf, _ := json.Marshal(entry.Data)
		return []byte(fmt.Sprintf("%v|%s|%s\n", entry.Time.Format(timeLayout), entry.Message, string(dataBuf))), nil
	} else {
		return []byte(fmt.Sprintf("%v|%s\n", entry.Time.Format(timeLayout), entry.Message)), nil
	}
}
