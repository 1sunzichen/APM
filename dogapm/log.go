package dogapm

import (
	"context"

	"github.com/sirupsen/logrus"
)

type log struct {
}

func init() {
	//将打印的格式设置为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

var Logger = &log{}

func (l *log) Info(ctx context.Context, action string, kv map[string]interface{}) {
	kv["action"] = action
	logrus.WithFields(kv).Info()

}

func (l *log) Error(ctx context.Context, action string, kv map[string]interface{}) {
	kv["action"] = action
	logrus.WithFields(kv).Error()
}

func (l *log) Debug(ctx context.Context, action string, kv map[string]interface{}) {
	kv["action"] = action
	logrus.WithFields(kv).Debug()
}

func (l *log) Warn(ctx context.Context, action string, kv map[string]interface{}) {
	kv["action"] = action
	logrus.WithFields(kv).Warn()
}
