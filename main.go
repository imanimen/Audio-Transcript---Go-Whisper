package main

import (
	"github.com/imanimen/Audio-Transcript---Go-Whisper/invokers"
	"github.com/imanimen/Audio-Transcript---Go-Whisper/providers"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(providers.NewConfig, providers.NewDatabase, providers.NewApi),
		fx.Invoke(invokers.ApiServer),
	).Run()
}
