package weatherleet

const forecastTemplate = `
Temp:        {{.Current.Main.Temp}} 
High:        {{.Current.Main.TempMax}} 
Low:         {{.Current.Main.TempMin}}

Feels like:  {{.Current.Main.FeelsLike}}
Humidity: 	 {{.Current.Main.Humidity}}
`
