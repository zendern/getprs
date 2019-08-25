package renderer

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/zendern/getprs/models"
)
import . "github.com/logrusorgru/aurora"

func RenderText(statuses []models.PRStatus) {
	for _, status := range statuses {
		fmt.Println(status.ApprovedStatus, Green(status.Title), "(", Bold(status.Username), ") - ", humanize.Time(status.TimeSinceOpened), "\n\t", Blue(status.PullRequestUrl))
	}
}