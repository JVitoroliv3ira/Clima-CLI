package views

import (
	"clima/src/types"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func Render(weather types.Weather) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cidade", "Status", "Temperatura", "Sensação Térmica", "Temperatura Máxima", "Temperatura Mínima"})
	table.Append([]string{
		fmt.Sprintf("%s, %s", weather.City, weather.Country.Name),
		weather.Info[0].Description,
		fmt.Sprintf("%.1f°C", weather.Climate.Temp),
		fmt.Sprintf("%.1f°C", weather.Climate.Sensation),
		fmt.Sprintf("%.1f°C", weather.Climate.Max),
		fmt.Sprintf("%.1f°C", weather.Climate.Min),
	})
	table.Render()
}
