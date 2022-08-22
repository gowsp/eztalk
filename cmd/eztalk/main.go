package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/eztalk/pkg/bot"
)

var (
	date    string
	version string
)
var (
	config string
	print  bool
)

// 解析参数
func parseArg() *bot.Bot {
	flag.BoolVar(&print, "v", false, "print version and exit")
	flag.StringVar(&config, "c", "", "Specifies an alternative per-user configuration file")
	flag.Parse()
	if print {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Release Date: %s\n", date)
		os.Exit(0)
	}
	if config == "" {
		fmt.Println("need config file")
		os.Exit(1)
	}
	bot, err := bot.NewByFile(config)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return bot
}

// 程序入口
func main() {
	bot := parseArg()
	go func() {
		if err := bot.Start(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	<-ctx.Done()
	log.Println("eztalk bot closed")
}
