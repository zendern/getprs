package renderer

import (
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"github.com/zendern/getprs/models"
	"os"
)
import . "github.com/logrusorgru/aurora"

func RenderTable(statuses []models.PRStatus){
	data := [][]string{}

	for _, status := range statuses {
		data = append(data, []string{
			status.ApprovedStatus, Green(status.Title).String(), Bold(status.Username).String(), humanize.Time(status.TimeSinceOpened), Blue(status.PullRequestUrl).String(),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Status", "Title", "User", "Opened", "PR Url"})
	table.SetAutoWrapText(false)

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}