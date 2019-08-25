package renderer

import (
	"github.com/olekukonko/tablewriter"
	"github.com/zendern/getprs/models"
	"os"
)
import . "github.com/logrusorgru/aurora"

func RenderTable(statuses []models.PRStatus){
	data := [][]string{}

	for _, status := range statuses {
		data = append(data, []string{
			status.ApprovedStatus, Green(status.Title).String(), Bold(status.Username).String(), Blue(status.PullRequestUrl).String(),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Approved Status", "Title", "User", "PR Url"})
	table.SetAutoWrapText(false)

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}