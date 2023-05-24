package config

import (
	"github.com/sakirsensoy/genv"
)

type cloudflareConf struct {
	TurnstileSecret string
	TurnstileUrl    string
}

var CLOUDFLARE = &cloudflareConf{
	TurnstileSecret: genv.Key("CLOUDFLARE_TURNSTILE_SECRET").Default("12345").String(),
	TurnstileUrl:    genv.Key("CLOUDFLARE_TURNSTILE_URL").Default("https://challenges.cloudflare.com/turnstile/v0/siteverify").String(),
}
