/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/gabeduke/weatherleet/pkg/weatherleet"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var weather *weatherleet.WeatherLeet

type conf struct {
	OwmApiKey string `env:"OWM_API_KEY,required"`
	Lat       string `env:"OWM_LAT" envDefault:"37.4978938"`
	Lon       string `env:"OWM_LON" envDefault:"-77.5504773"`
}

var rootCmd = &cobra.Command{
	Use:   "weatherleet",
	Short: "For all your weather needs..",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	var err error

	c := conf{}
	if err := env.Parse(&c); err != nil {
		log.Fatal(err)
	}

	weather, err = weatherleet.New(c.OwmApiKey, c.Lat, c.Lon)
	if err != nil {
		log.Fatal(err)
	}
}
