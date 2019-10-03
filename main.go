package main

import (
	"github.com/zendern/getprs/models"
	"github.com/zendern/getprs/renderer"
	"golang.org/x/oauth2"
	"sort"
	"strconv"
)
import "github.com/google/go-github/github"
import (
	"fmt"
	"golang.org/x/net/context"
	"os"
	"strings"
)
import . "github.com/logrusorgru/aurora"

var options = &github.ListOptions{PerPage: 1000}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Arguments required. <personal access token> <organization> <team name> <renderer [text, json, table] optional>")
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

	var renderStr string
	if len(os.Args) < 5 || strings.TrimSpace(os.Args[4]) == "" {
		renderStr = "table"
	} else {
		renderStr = strings.TrimSpace(os.Args[4])
	}
	renderFn, ok := renderer.Renderers[renderStr]
	if !ok {
		fmt.Println("Renderer must be one of the allowed values. [table,text,json]")
		os.Exit(1)
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

	org := getOrgByName(ctx, client, orgName)
	foundTeam := getTeam(ctx, client, org, teamName)
	teamMembers := getAllTeamMembers(ctx, client, foundTeam)
	issues := getAllOpenPRs(ctx, client, org, foundTeam, teamMembers)
	statuses := getPRStatuses(ctx, client, org, issues)

	sort.Sort(models.ByStatus(statuses))

	renderFn(statuses)
}

func getOrgByName(ctx context.Context, client *github.Client, orgName string) *github.Organization {
	fmt.Println(">>> FINDING ORG BY NAME : ", Bold(orgName))
	org, _, err := client.Organizations.Get(ctx, orgName)
	if err != nil {
		fmt.Println(">>> Failed to find organization with name - " + orgName + "<<< : " + err.Error())
		os.Exit(1)
	}
	return org
}

func getPRStatuses(ctx context.Context, client *github.Client, org *github.Organization, issues []github.Issue) []models.PRStatus {
	statuses := make([]models.PRStatus, 0)
	for _, issue := range issues {
		positionOfLastSlash := strings.LastIndex(*issue.RepositoryURL, "/")
		repoUrl := *issue.RepositoryURL
		repoName := repoUrl[positionOfLastSlash+1 : len(repoUrl)]
		prReviews, _, err := client.PullRequests.ListReviews(ctx, *org.Name, repoName, *issue.Number, options)
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
		} else {
			uiApprovedState = "\u274C"
		}

		statuses = append(statuses, models.PRStatus{
			Username:        *issue.User.Login,
			Title:           *issue.Title,
			ApprovedStatus:  uiApprovedState,
			PullRequestUrl:  *issue.HTMLURL,
			TimeSinceOpened: *issue.CreatedAt,
		})
	}
	return statuses
}

func getTeam(ctx context.Context, client *github.Client, org *github.Organization, teamName string) *github.Team {
	fmt.Println(">>> GETTING ALL TEAMS FOR ORG: ", Bold(*org.Name))
	teams, _, err := client.Teams.ListTeams(ctx, *org.Name, options)
	if err != nil {
		fmt.Println(">>> Failed to find teams for org with name - " + *org.Name + "<<< : " + err.Error())
		os.Exit(1)
	}
	var foundTeam *github.Team
	for _, team := range teams {
		if *team.Name == teamName || *team.Slug == teamName {
			foundTeam = team
		}
	}
	if foundTeam == nil {
		fmt.Println(">>> Failed to find team with name - " + teamName)
		os.Exit(1)
	}
	return foundTeam
}

func getAllOpenPRs(ctx context.Context, client *github.Client, org *github.Organization, foundTeam *github.Team, teamMembers []*github.User) []github.Issue {
	fmt.Println(">>> FINDING ALL OPEN PRS FOR TEAM : ", Bold(*foundTeam.Name))
	fmt.Println("\n")
	searchOpts := &github.SearchOptions{ListOptions: *options}
	q := fmt.Sprintf("org:%s is:open is:pr", *org.Name)
	for _, member := range teamMembers {
		q += " author:" + *member.Login
	}
	issues, _, err := client.Search.Issues(ctx, q, searchOpts)
	if err != nil {
		fmt.Println(">>> Failed to find issues for query - " + q + "<<< : " + err.Error())
		os.Exit(1)
	}

	/*
		For whatever reason even searching for is:open might return close PR ¯\_(ツ)_/¯ So we loop and exclude them to make it really what we want
	*/
	actualOpenedIssues := make([]github.Issue, 0)
	for _, issue := range issues.Issues {
		if *issue.State == "open" {
			actualOpenedIssues = append(actualOpenedIssues, issue)
		}
	}
	return actualOpenedIssues
}

func getAllTeamMembers(ctx context.Context, client *github.Client, foundTeam *github.Team) []*github.User {
	fmt.Println(">>> GETTING MEMBERS ON TEAM : ", Bold(*foundTeam.Name))
	teamMemberOpts := &github.TeamListTeamMembersOptions{ListOptions: *options}
	teamMembers, _, err := client.Teams.ListTeamMembers(ctx, *foundTeam.ID, teamMemberOpts)
	if err != nil {
		fmt.Println(">>> Failed to find team members for team name - " + *foundTeam.Slug + "<<< : " + err.Error())
		os.Exit(1)
	}
	return teamMembers
}

func Any(vs []*github.PullRequestReview, f func(review *github.PullRequestReview) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
