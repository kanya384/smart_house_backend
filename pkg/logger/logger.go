package logger

import (
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	formatter "github.com/fabienm/go-logrus-formatters"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(msg string, params interface{})
	Info(msg string, params interface{})
	Warn(msg string, params interface{})
	Error(msg string, err error, params interface{})
	Panic(msg string, err error, params interface{})
}

type logger struct {
	log         *logrus.Logger
	serviceName string
}

func NewLogger(serviceName string, logLevel uint32, logFile string) (Logger, error) {
	log := logrus.New()
	log.SetLevel(logrus.Level(logLevel))
	gelFmt := formatter.NewGelf(serviceName)
	runtimeFormatter := &runtime.Formatter{ChildFormatter: gelFmt}
	log.SetFormatter(runtimeFormatter)
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)
	//hook := graylog.NewGraylogHook(greyLogHost, map[string]interface{}{})
	//log.AddHook(hook)
	defer file.Close()

	return &logger{
		log: log,
	}, nil
}

func (lg *logger) Debug(msg string, params interface{}) {
	logger := lg.log
	logger.WithField("service_name", lg.serviceName)
	for key, param := range params.(map[string]interface{}) {
		logger.WithField(key, param)
	}
	logger.Debug(msg)
}

func (lg *logger) Info(msg string, params interface{}) {
	logger := lg.log
	logger.WithField("service_name", lg.serviceName)
	if params != nil {
		for key, param := range params.(map[string]interface{}) {
			logger.WithField(key, param)
		}
	}
	logger.Info(msg)
}

func (lg *logger) Warn(msg string, params interface{}) {
	logger := lg.log
	logger.WithField("service_name", lg.serviceName)
	if params != nil {
		for key, param := range params.(map[string]interface{}) {
			logger.WithField(key, param)
		}
	}
	logger.Warn(msg)
}

func (lg *logger) Error(msg string, err error, params interface{}) {
	logger := lg.log
	logger.WithField("service_name", lg.serviceName)
	if params != nil {
		for key, param := range params.(map[string]interface{}) {
			logger.WithField(key, param)
		}
	}
	logger.WithError(errors.WithStack(err)).Error(err)
}

func (lg *logger) Panic(msg string, err error, params interface{}) {
	logger := lg.log
	logger.WithField("service_name", lg.serviceName)
	if params != nil {
		for key, param := range params.(map[string]interface{}) {
			logger.WithField(key, param)
		}
	}
	logger.WithError(errors.WithStack(err)).Panic(err)
}
