package internal

import (
	"context"
	"github.com/hashicorp/go-tfe"
	"github.com/peytoncasper/tfe-usage-stats/log"
	"go.uber.org/zap"
	"time"
)

func GetWorkspaces(client *tfe.Client, organizations []*tfe.Organization) ([]*tfe.Workspace, error) {
	workspaces := make([]*tfe.Workspace, 0)

	for _, org := range organizations {
		orgWorkspaces, err := getOrganizationWorkspaces(client, org)
		workspaces = append(workspaces, orgWorkspaces...)

		if err != nil {
			return orgWorkspaces, err
		}

	}

	return workspaces, nil
}

func getOrganizationWorkspaces(client *tfe.Client, organization *tfe.Organization) ([]*tfe.Workspace, error) {
	workspaces := make([]*tfe.Workspace, 0)

	currentPage := 0
	totalPages := 1
	pageSize := 20

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
		log.Debug("added workspaces", zap.Int("count", len(workspaces)))
		log.Debug("page", zap.Int("current", currentPage), zap.Int("total", totalPages))

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
