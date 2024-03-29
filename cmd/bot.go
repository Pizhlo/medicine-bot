package bot

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/Pizhlo/medicine-bot/internal/config"
	"github.com/Pizhlo/medicine-bot/internal/controller"
	"github.com/Pizhlo/medicine-bot/internal/server"
	"github.com/Pizhlo/medicine-bot/internal/service/user"
	user_storage "github.com/Pizhlo/medicine-bot/internal/storage/user"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func Start(confName, path string) {
	logrus.SetLevel(logrus.DebugLevel)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf, err := config.LoadConfig(confName, path)
	if err != nil {
		logrus.Fatalf("unable to load config: %v", err)
	}

	dbUser, err := user_storage.New(conf.DBAddress)
	if err != nil {
		logrus.Fatalf("unable to connect DB: %v", err)
	}

	userSrv := user.New(dbUser)

	controller := controller.New(userSrv)

	bot, err := tele.NewBot(tele.Settings{
		Token:  conf.Token,
		Poller: &tele.LongPoller{Timeout: conf.Timeout},
	})
	if err != nil {
		logrus.Fatalf("cannot create a bot: %v", err)
	}

	server := server.New(bot, controller)

	logrus.Debug("starting server...")

	go func() {
		bot.Start()
	}()

	server.Start(ctx)

	notifyCtx, notify := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer notify()

	logrus.Info("started")

	<-notifyCtx.Done()
	logrus.Info("shutdown")

	go func() {
		bot.Stop()
	}()

	notify()
}
