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
	"github.com/gabeduke/weatherleet/functions"
	"github.com/go-zepto/zepto"
	"github.com/gorilla/mux"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start weather server",
	Run: func(cmd *cobra.Command, args []string) {
		z := zepto.NewZepto(
			zepto.Name("leetserve.weather"),
			zepto.Version("latest"),
		)

		r := mux.NewRouter()

		w := functions.New(weather)

		r.HandleFunc("/current", w.CurrentWeather)
		r.HandleFunc("/minute", w.MinutelyWeather)
		r.HandleFunc("/hour", w.HourlyWeather)
		r.HandleFunc("/day", w.DailyWeather)

		// Setup HTTP Server
		z.SetupHTTP("0.0.0.0:8080", r)

		z.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
