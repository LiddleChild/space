package config

import "github.com/spf13/cobra"

var AppConfig *Config

func init() {
	var err error
	AppConfig, err = Load()
	cobra.CheckErr(err)
}
