package internal

import (
	"context"
	"github.com/hashicorp/go-tfe"
	"github.com/peytoncasper/tfe-usage-stats/log"
	"time"
)

func GetWorkspaces(client *tfe.Client, organizations []*tfe.Organization) ([]*tfe.Workspace, error) {
	workspaces := make([]*tfe.Workspace, 0)

	for _, org := range organizations {
		orgWorkspaces, err := getOrganizationWorkspaces(client, org)
		if err != nil {
			return workspaces, err
		}

		workspaces = append(workspaces, orgWorkspaces...)
	}

	return workspaces, nil
}

func getOrganizationWorkspaces(client *tfe.Client, organization *tfe.Organization) ([]*tfe.Workspace, error) {
	workspaces := make([]*tfe.Workspace, 0)

	currentPage := 0
	totalPages := 1
	pageSize := 10

	for currentPage < totalPages {
		workspacePage, err := getWorkspacePage(client, organization.Name, tfe.WorkspaceListOptions{
			ListOptions: tfe.ListOptions{
				PageNumber: currentPage,
				PageSize:   pageSize,
			},
		})

		if err != nil {
			log.Error("error getting workspace page, retrying in 10 seconds")
			time.Sleep(10 * time.Second)
			workspacePage, err = getWorkspacePage(client, organization.Name, tfe.WorkspaceListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: currentPage,
					PageSize:   pageSize,
				},
			})

			if err != nil {
				return workspaces, err
			}
		}

		workspaces = append(workspaces, workspacePage.Items...)

		totalPages = workspacePage.TotalPages
		currentPage++
	}

	return workspaces, nil
}

func getWorkspacePage(client *tfe.Client, organizationName string, options tfe.WorkspaceListOptions) (*tfe.WorkspaceList, error) {
	workspaces, err := client.Workspaces.List(context.Background(), organizationName, options)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}
