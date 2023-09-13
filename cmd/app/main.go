package main

import (
	"github.com/GrishaSkurikhin/EffectiveMobileTesting/internal/config"
)

const configPath = "config/.env"

func init() {
	config.MustLoad(configPath)
}

func main() {

}
