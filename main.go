package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"clima/src/services"
	"clima/src/views"
)

var city string

var rootCmd = &cobra.Command{
	Use:   "clima",
	Short: `Clima CLI é uma ferramenta de linha de comando (CLI) simples para obter informações climáticas de uma cidade específica. Utilizando a API da OpenWeather, a Clima CLI fornece dados como temperatura atual, sensação térmica, temperatura máxima e mínima, e descrição das condições climáticas.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENWEATHER_API_KEY")
		if apiKey == "" {
			fmt.Println("Chave de API não definida. Por favor, defina a variável de ambiente OPENWEATHER_API_KEY.")
			os.Exit(1)
		}
		data, err := services.Fetch(
			fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&lang=pt_br&appid=%s", city, apiKey),
		)

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		views.Render(data)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&city, "city", "c", "Natal", "city")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
