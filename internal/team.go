package internal

import (
	"context"
	"github.com/hashicorp/go-tfe"
)

func GetTeams(client *tfe.Client, organizations []*tfe.Organization) ([]*tfe.Team, error) {
	teams := make([]*tfe.Team, 0)

	for _, org := range organizations {
		orgTeams, err := getOrganizationTeams(client, org)
		if err != nil { return nil, err }

		teams = append(teams, orgTeams...)
	}

	return teams, nil
}

func getOrganizationTeams(client *tfe.Client, organization *tfe.Organization) ([]*tfe.Team, error) {
	teams := make([]*tfe.Team, 0)

	currentPage := 0
	totalPages := 1
	pageSize := 10

	for currentPage < totalPages {
		teamPage, err := getTeamPage(client, organization.Name, tfe.TeamListOptions{
			ListOptions: tfe.ListOptions{
				PageNumber: currentPage,
				PageSize:   pageSize,
			},
		})

		if err != nil { return nil, err }

		teams = append(teams, teamPage.Items...)

		totalPages = teamPage.TotalPages
		currentPage++
	}

	return teams, nil
}

func getTeamPage(client *tfe.Client, organizationName string, options tfe.TeamListOptions) (*tfe.TeamList, error) {
	teams, err := client.Teams.List(context.Background(), organizationName, options)
	if err != nil {
		return nil, err
	}

	return teams, nil
}
