# Gets PRs by Org and Team
[![Build Status](https://travis-ci.com/zendern/getprs.svg?branch=master)](https://travis-ci.com/zendern/getprs)
Most of the time you want to be able to easily look at the PRs for your team so that they can be reviewed on an occasional basis. This application simplifies that by doing all the leg work for you.

# How to Install
Download the executable found in this repo

* MacOSX : [getprs](https://github.com/zendern/getprs/releases/latest/download/getprs-darwin-amd64)
* Linux : [getprs](https://github.com/zendern/getprs/releases/latest/download/getprs-linux-amd64)
* Windows (32) : [getprs.exe](https://github.com/zendern/getprs/releases/latest/download/getprs-windows-386.exe)
* Windows (64) : [getprs.exe](https://github.com/zendern/getprs/releases/latest/download/getprs-windows-amd64.exe)

Or clone the repo and build it yourself for whatever platform you need.

## How to build it
Pre-Req: 
Have go 1.11 or greater installed
1. Clone the repo
2. Execute the following build script ./build/build.sh

# How to run it

Mac OSX
```
./getprs-darwin-amd64 <github api token> <organization> <team name> <optional output format>
```

Windows
```
getprs-windows-386.exe <github api token> <organization> <team name> <optional output format>
getprs-windows-amd64.exe <github api token> <organization> <team name> <optional output format>
```

## Where do i get this github api token thing??!?!?

See [here](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) for a guide on how to do that 

# Screenshot

Table output (default)
```
>>> FINDING ORG BY NAME :  cahcommercial
>>> GETTING ALL TEAMS FOR ORG:  cahcommercial
>>> GETTING MEMBERS ON TEAM :  Squid Squad
>>> FINDING ALL OPEN PRS FOR TEAM :  Squid Squad


+--------+------------------------------------------------------------------------+---------------+-------------+----------------------------------------------------------------------+
| STATUS |                                 TITLE                                  |     USER      |   OPENED    |                                PR URL                                |
+--------+------------------------------------------------------------------------+---------------+-------------+----------------------------------------------------------------------+
| ✅     | CONNECT-7342 correct URL for local feature flags                       | cah-erinblake | 1 week ago  | https://github.com/cahcommercial/connect-services/pull/240           |
| ✅     | CONNECT-7313 - added pact provider tests                               | cah-markrave  | 2 weeks ago | https://github.com/cahcommercial/connect-reporting-services/pull/243 |
| ✅     | CONNECT-7313 - First attempt at pact testing.  Still needs work in th… | cah-markrave  | 2 weeks ago | https://github.com/cahcommercial/connect-reporting-client/pull/138   |
+--------+------------------------------------------------------------------------+---------------+-------------+----------------------------------------------------------------------+
+--------+------------------------------------------------------------------+------------------+--------------+------------------------------------------------------------------------------+
| STATUS |                              TITLE                               |       USER       |    OPENED    |                                    PR URL                                    |
+--------+------------------------------------------------------------------+------------------+--------------+------------------------------------------------------------------------------+
| ❌     | Connect 7354                                                     | cah-erinblake    | 4 days ago   | https://github.com/cahcommercial/connect-data-generators/pull/88             |
| ❌     | Getting goreleaser and a concourse pipeline in place             | cah-nathanzender | 2 weeks ago  | https://github.com/cahcommercial/fusecli/pull/4                              |
| ❌     | Deploy keys yo                                                   | cah-nathanzender | 3 weeks ago  | https://github.com/cahcommercial/fusecli/pull/3                              |
| ❌     | Add an aws token generator                                       | cah-nathanzender | 3 weeks ago  | https://github.com/cahcommercial/fusecli/pull/2                              |
| ❌     | CONNECT-5980: ActiveOpportunityConsumer setup                    | cah-ryanhipps    | 1 month ago  | https://github.com/cahcommercial/connect-services/pull/220                   |
| ❌     | Proposal on how to go about fixing Concourse/Github webhooks     | cah-nathanzender | 2 months ago | https://github.com/cahcommercial/concourse-github-webhook-broadcaster/pull/1 |
| ❌     | INTERNAL Trying out the webhook stuff                            | cah-nathanzender | 3 months ago | https://github.com/cahcommercial/connect-reporting-services/pull/149         |
| ❌     | INTERNAL Update connection details for kafka upgrade             | cah-nathanzender | 4 months ago | https://github.com/cahcommercial/connect-analytics-infrastructure/pull/126   |
| ❌     | CONNECT-5923 - APM RUM stuff                                     | cah-timhuddle    | 4 months ago | https://github.com/cahcommercial/connect-reporting-client/pull/78            |
| ❌     | INTERNAL - parallel at feature level                             | cah-tylerrasor   | 7 months ago | https://github.com/cahcommercial/connect-reporting-client/pull/25            |
| ❌     | INTERNAL - parallel at scenario level                            | cah-tylerrasor   | 7 months ago | https://github.com/cahcommercial/connect-reporting-client/pull/24            |
| ❌     | INTERNAL - use current UTC offset to advance debezium timestamps | cah-tylerrasor   | 8 months ago | https://github.com/cahcommercial/connect-analytics-infrastructure/pull/71    |
+--------+------------------------------------------------------------------+------------------+--------------+------------------------------------------------------------------------------+
```

JSON output
```
>>> FINDING ORG BY NAME :  cahcommercial
>>> GETTING ALL TEAMS FOR ORG:  cahcommercial
>>> GETTING MEMBERS ON TEAM :  Squid Squad
>>> FINDING ALL OPEN PRS FOR TEAM :  Squid Squad


[
	{
		"Username": "cah-erinblake",
		"Title": "Connect 7354",
		"Approved": false,
		"ApprovedStatus": "❌",
		"PullRequestUrl": "https://github.com/cahcommercial/connect-data-generators/pull/88",
		"TimeSinceOpened": "2019-12-16T18:46:01Z"
	},
	{
		"Username": "cah-nathanzender",
		"Title": "Getting goreleaser and a concourse pipeline in place",
		"Approved": false,
		"ApprovedStatus": "❌",
		"PullRequestUrl": "https://github.com/cahcommercial/fusecli/pull/4",
		"TimeSinceOpened": "2019-11-30T19:46:43Z"
	},
	{
		"Username": "cah-nathanzender",
		"Title": "Deploy keys yo",
		"Approved": false,
		"ApprovedStatus": "❌",
		"PullRequestUrl": "https://github.com/cahcommercial/fusecli/pull/3",
		"TimeSinceOpened": "2019-11-27T20:51:27Z"
	},
	{
		"Username": "cah-nathanzender",
		"Title": "Add an aws token generator",
		"Approved": false,
		"ApprovedStatus": "❌",
		"PullRequestUrl": "https://github.com/cahcommercial/fusecli/pull/2",
		"TimeSinceOpened": "2019-11-25T03:48:42Z"
	},
	{
		"Username": "cah-ryanhipps",
		"Title": "CONNECT-5980: ActiveOpportunityConsumer setup",
		"Approved": false,
		"ApprovedStatus": "❌",
		"PullRequestUrl": "https://github.com/cahcommercial/connect-services/pull/220",
		"TimeSinceOpened": "2019-11-04T15:07:32Z"
	},
	{
		"Username": "cah-nathanzender",
		"Title": "Proposal on how to go about fixing Concourse/Github webhooks",
		"Approved": false,
		"ApprovedStatus": "❌",
		"PullRequestUrl": "https://github.com/cahcommercial/concourse-github-webhook-broadcaster/pull/1",
		"TimeSinceOpened": "2019-09-26T01:48:09Z"
```

Text output
```
>>> FINDING ORG BY NAME :  cahcommercial
>>> GETTING ALL TEAMS FOR ORG:  cahcommercial
>>> GETTING MEMBERS ON TEAM :  Squid Squad
>>> FINDING ALL OPEN PRS FOR TEAM :  Squid Squad


❌ Connect 7354 ( cah-erinblake ) -  4 days ago
	 https://github.com/cahcommercial/connect-data-generators/pull/88
❌ Getting goreleaser and a concourse pipeline in place ( cah-nathanzender ) -  2 weeks ago
	 https://github.com/cahcommercial/fusecli/pull/4
❌ Deploy keys yo ( cah-nathanzender ) -  3 weeks ago
	 https://github.com/cahcommercial/fusecli/pull/3
❌ Add an aws token generator ( cah-nathanzender ) -  3 weeks ago
	 https://github.com/cahcommercial/fusecli/pull/2
❌ CONNECT-5980: ActiveOpportunityConsumer setup ( cah-ryanhipps ) -  1 month ago
	 https://github.com/cahcommercial/connect-services/pull/220
❌ Proposal on how to go about fixing Concourse/Github webhooks ( cah-nathanzender ) -  2 months ago
	 https://github.com/cahcommercial/concourse-github-webhook-broadcaster/pull/1
❌ INTERNAL Trying out the webhook stuff ( cah-nathanzender ) -  3 months ago
	 https://github.com/cahcommercial/connect-reporting-services/pull/149
❌ INTERNAL Update connection details for kafka upgrade ( cah-nathanzender ) -  4 months ago
	 https://github.com/cahcommercial/connect-analytics-infrastructure/pull/126
❌ CONNECT-5923 - APM RUM stuff ( cah-timhuddle ) -  4 months ago
	 https://github.com/cahcommercial/connect-reporting-client/pull/78
❌ INTERNAL - parallel at feature level ( cah-tylerrasor ) -  7 months ago
	 https://github.com/cahcommercial/connect-reporting-client/pull/25
❌ INTERNAL - parallel at scenario level ( cah-tylerrasor ) -  7 months ago
	 https://github.com/cahcommercial/connect-reporting-client/pull/24
❌ INTERNAL - use current UTC offset to advance debezium timestamps ( cah-tylerrasor ) -  8 months ago
	 https://github.com/cahcommercial/connect-analytics-infrastructure/pull/71
✅ CONNECT-7342 correct URL for local feature flags ( cah-erinblake ) -  1 week ago
	 https://github.com/cahcommercial/connect-services/pull/240
✅ CONNECT-7313 - added pact provider tests ( cah-markrave ) -  2 weeks ago
	 https://github.com/cahcommercial/connect-reporting-services/pull/243
✅ CONNECT-7313 - First attempt at pact testing.  Still needs work in th… ( cah-markrave ) -  2 weeks ago
	 https://github.com/cahcommercial/connect-reporting-client/pull/138
```

