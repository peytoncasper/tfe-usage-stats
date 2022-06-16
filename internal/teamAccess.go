package internal

import (
	"context"
	"github.com/hashicorp/go-tfe"
)

func GetTeamAccessRelationships(client *tfe.Client, workspaces []*tfe.Workspace) ([]*tfe.TeamAccess, error) {
	teamAccess := []*tfe.TeamAccess{}

	for _, workspace := range workspaces {
		relations, err := getWorkspaceTeamAccessRelationships(client, workspace)
		if err != nil {
			return nil, err
		}

		teamAccess = append(teamAccess, relations...)
	}

	return teamAccess, nil
}

func getWorkspaceTeamAccessRelationships(client *tfe.Client, workspace *tfe.Workspace) ([]*tfe.TeamAccess, error) {

	currentPage := 0
	totalPages := 1
	pageSize := 30

	relationships := []*tfe.TeamAccess{}

	for currentPage < totalPages {
		relationshipsPage, err := getTeamAccessPage(client, workspace.ID, tfe.TeamAccessListOptions{
			ListOptions: tfe.ListOptions{
				PageNumber: currentPage,
				PageSize:   pageSize,
			},
			WorkspaceID: &workspace.ID,
		})

		if err != nil {
			return relationships, err
		}

		for _, relation := range relationshipsPage.Items {
			relation.Team, err = client.Teams.Read(context.Background(), relation.Team.ID)
			relation.Workspace, err = client.Workspaces.ReadByID(context.Background(), relation.Workspace.ID)

			relationships = append(relationships, relation)
		}

		totalPages = relationshipsPage.TotalPages
		currentPage++
	}

	return relationships, nil
}

func getTeamAccessPage(client *tfe.Client, workspaceId string, options tfe.TeamAccessListOptions) (*tfe.TeamAccessList, error) {
	relationships, err := client.TeamAccess.List(context.Background(), options)
	if err != nil {
		return nil, err
	}

	return relationships, nil
}
