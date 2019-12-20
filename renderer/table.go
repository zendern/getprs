package renderer

import (
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"github.com/zendern/getprs/models"
	"os"
	"sort"
)
import . "github.com/logrusorgru/aurora"

func RenderTable(statuses []models.PRStatus){
	approvedData := []models.PRStatus{}
	needsWorkData := []models.PRStatus{}

	for _, status := range statuses {
		if status.Approved {
			approvedData = append(approvedData, status)
		}else{
			needsWorkData = append(needsWorkData, status)
		}
	}

	generateTableFromData(approvedData)
	generateTableFromData(needsWorkData)
}

func generateTableFromData(data []models.PRStatus) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Status", "Title", "User", "Opened", "PR Url"})
	table.SetAutoWrapText(false)

	sort.Sort(models.ByStatusAndTime(data))

	for _, status := range data {
		table.Append([]string{
			status.ApprovedStatus, Green(status.Title).String(), Bold(status.Username).String(), humanize.Time(status.TimeSinceOpened), Blue(status.PullRequestUrl).String(),
		})
	}
	table.Render()
	// Send output
}