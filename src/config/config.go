package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port         int
	DBdriver     string
	DBhost       string
	DBname       string
	DBuser       string
	DBpassword   string
	StripeKey    string
	PK           string
	EPS          string
	Recaptcha    string
	GmailAddr    string
	GmailUser    string
	GmailPass    string
	AdminPass    string
	AdminToken   string
	SendMailAddr string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Port:         cfg.Section("web").Key("port").MustInt(),
		DBdriver:     cfg.Section("db").Key("driver").String(),
		DBhost:       cfg.Section("db").Key("db_host").String(),
		DBname:       cfg.Section("db").Key("name").String(),
		DBuser:       cfg.Section("db").Key("user").String(),
		DBpassword:   cfg.Section("db").Key("password").String(),
		StripeKey:    cfg.Section("stripe").Key("stripe_key").String(),
		PK:           cfg.Section("stripe").Key("publish_key").String(),
		EPS:          cfg.Section("stripe").Key("eps").String(),
		Recaptcha:    cfg.Section("recaptcha").Key("pk").String(),
		GmailAddr:    cfg.Section("email").Key("addr").String(),
		GmailUser:    cfg.Section("email").Key("username").String(),
		GmailPass:    cfg.Section("email").Key("password").String(),
		AdminPass:    cfg.Section("admin").Key("pass").String(),
		AdminToken:   cfg.Section("admin").Key("token").String(),
		SendMailAddr: cfg.Section("email").Key("sendmailaddr").String(),
	}
}
