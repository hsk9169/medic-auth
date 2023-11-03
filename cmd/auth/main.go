package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	_ "gopkg.in/yaml.v2"

	"github.com/medic-basic/auth/pkg/domain"
	"github.com/medic-basic/auth/pkg/external/cache"
	"github.com/medic-basic/auth/pkg/server"
)

const version = "0.0.1"
const defaultLogDir = "logs"
const defaultLogFileName = "access.log"

var (
	confFilePath string
	cfg          *viper.Viper
	srv          *server.Server
)

func parseCmdLineInput() error {
	flag.StringVar(&confFilePath, "config", "", "input config filepath")
	flag.Parse()

	if confFilePath == "" {
		return errors.New("empty config filepath")
	}

	if _, err := os.Stat(confFilePath); os.IsNotExist(err) {
		return errors.Errorf("invalid input config filepath(%s)", confFilePath)
	}

	return nil
}

func getLogger(logFileName string) *lumberjack.Logger {
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})

	os.Mkdir(defaultLogDir, os.ModePerm)
	loggerForRotate := &lumberjack.Logger{
		Filename:   defaultLogDir + "/" + logFileName,
		MaxSize:    64,
		MaxBackups: 5,
		MaxAge:     1,
		LocalTime:  true,
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for {
			<-c
			loggerForRotate.Rotate()
		}
	}()

	return loggerForRotate
}

func init() {
	if err := parseCmdLineInput(); err != nil {
		panic(err)
	}
	if err := load(confFilePath); err != nil {
		panic(err)
	}

	logrus.SetOutput(io.MultiWriter(getLogger(defaultLogFileName), os.Stdout))

	//redis.Init()
	cache.Init()

}

func main() {
	defer closeAll()
	srv = server.NewServer(cfg.GetInt("server.port"), version)

	go func(srv *server.Server) {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(errors.Wrap(err, "server listen failed"))
		}
	}(srv)

	logrus.Infof("start auth server(version:%s) port %d", version, cfg.GetInt("server.port"))

	//domain.Init(domain.Config{RedisClient: &redis.CacheImpl})
	domain.Init(domain.Config{CacheClient: &cache.CacheImpl})
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func closeAll() {
	logrus.Info("closing auth server...")
	srv.Close()
}

func load(configFilePath string) error {
	cfg = viper.New()
	cfg.SetConfigFile(configFilePath)

	if err := cfg.ReadInConfig(); err != nil {
		panic(err)
	}

	return nil
}
