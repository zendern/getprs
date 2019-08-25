package main

import (
	"github.com/zendern/getprs/models"
	"github.com/zendern/getprs/renderer"
	"golang.org/x/oauth2"
	"strconv"
)
import "github.com/google/go-github/github"
import (
	"golang.org/x/net/context"
	"fmt"
	"os"
	"strings"
)
import . "github.com/logrusorgru/aurora"

func main() {
	if len(os.Args) < 4  {
		fmt.Println("Arguments required. <personal access token> <organization> <team name> <renderer [text, json] optional>")
		os.Exit(1)
	}
	if strings.Trim(os.Args[1], " ") == "" {
		fmt.Println("Personal Access Token must not be blank")
		os.Exit(1)
	}
	if strings.Trim(os.Args[2], " ") == "" {
		fmt.Println("Organization name must not be blank")
		os.Exit(1)
	}
	if strings.Trim(os.Args[3], " ") == "" {
		fmt.Println("Team name must not be blank")
		os.Exit(1)
	}

	fmt.Println(os.Args)
	var renderType string
	if len(os.Args) == 5 {
		trimmedRenderType := strings.Trim(os.Args[4], " ")
		if 	trimmedRenderType == "" {
			renderType = "text"
		}else{
			if trimmedRenderType == "text" || trimmedRenderType == "json"{
				renderType = os.Args[4]
			}else{
				fmt.Println("Renderer must be one of the allowed values. [txt or json]")
				os.Exit(1)
			}
		}
	}else{
		renderType = "text"
	}

	accessToken := os.Args[1]
	orgName := os.Args[2]
	teamName := os.Args[3]

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	fmt.Println(">>> FINDING ORG BY NAME : ", Bold(orgName))
	org, _, err := client.Organizations.Get(ctx, orgName)
	if err != nil {
		fmt.Println(">>> Failed to find organization with name - " + orgName + "<<< : " + err.Error())
		os.Exit(1)
	}

	fmt.Println(">>> GETTING ALL TEAMS FOR ORG: ", Bold(orgName))
	options := &github.ListOptions{PerPage: 1000}
	teams, _, err := client.Teams.ListTeams(ctx, *org.Name, options)
	if err != nil {
		fmt.Println(">>> Failed to find teams for org with name - " + orgName + "<<< : " + err.Error())
		os.Exit(1)
	}

	var foundTeam *github.Team;
	for _, team := range teams {
		if *team.Name == teamName || *team.Slug == teamName {
			foundTeam = team
		}
	}

	if foundTeam == nil {
		fmt.Println(">>> Failed to find team with name - " + teamName)
		os.Exit(1)
	}

	fmt.Println(">>> GETTING MEMBERS ON TEAM : ", Bold(teamName))
	teamMemberOpts := &github.TeamListTeamMembersOptions{ListOptions: *options}
	teamMembers, _, err := client.Teams.ListTeamMembers(ctx, *foundTeam.ID, teamMemberOpts)
	if err != nil {
		fmt.Println(">>> Failed to find team members for team name - " + *foundTeam.Slug + "<<< : " + err.Error())
		os.Exit(1)
	}

	fmt.Println(">>> FINDING ALL OPEN PRS FOR TEAM : ", Bold(teamName))
	fmt.Println("\n")
	searchOpts := &github.SearchOptions{ListOptions: *options}
	q := fmt.Sprintf("org:%s is:open is:pr", orgName)
	for _, member := range teamMembers {
		q += " author:" + *member.Login
	}
	issues, _, err := client.Search.Issues(ctx, q, searchOpts)
	if err != nil {
		fmt.Println(">>> Failed to find issues for query - " + q + "<<< : " + err.Error())
		os.Exit(1)
	}

	statuses := make([]models.PRStatus, 0)
	for _, issue := range issues.Issues {
		positionOfLastSlash := strings.LastIndex(*issue.RepositoryURL, "/")
		repoUrl := *issue.RepositoryURL;
		repoName := repoUrl[positionOfLastSlash+1 : len(repoUrl)]
		prReviews, _, err := client.PullRequests.ListReviews(ctx, orgName, repoName, *issue.Number, options)
		if err != nil {
			fmt.Println(">>> Failed to find PR reviews for issue - " + strconv.Itoa(*issue.Number) + "<<< : " + err.Error())
			os.Exit(1)
		}
		hasBeenApproved := Any(prReviews, func(s *github.PullRequestReview) bool {
			return *s.State == "APPROVED"
		})
		var uiApprovedState string
		if hasBeenApproved {
			uiApprovedState = "\u2705"
		} else{
			uiApprovedState = "\u274C"
		}

		statuses = append(statuses, models.PRStatus{
			Username: *issue.User.Login,
			Title: *issue.Title,
			ApprovedStatus: uiApprovedState,
			PullRequestUrl: *issue.HTMLURL,
		})
	}

	if renderType == "json" {
		renderer.RenderJson(statuses)
	} else if renderType == "text" {
		renderer.RenderText(statuses)
	}
}

func Any(vs []*github.PullRequestReview, f func(review *github.PullRequestReview) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
