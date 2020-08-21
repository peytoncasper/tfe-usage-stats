package internal

import (
	"context"
	tfe "github.com/hashicorp/go-tfe"
	"time"
)

func GetRuns(client *tfe.Client, workspaces []*tfe.Workspace) (map[string][]*tfe.Run, error) {
	runs := map[string][]*tfe.Run{}

	now := time.Now()
	for i := 0; i <= 5; i++ {
		runs[now.AddDate(0, i * -1, (now.Day() - 1) * -1).Format("2006-01-02")] = make([]*tfe.Run, 0)
	}


	for _, workspace := range workspaces {
		err := getWorkspaceRuns(client, workspace, runs)
		if err != nil { return nil, err }

		//print(workspaceRuns)
	//
	//	runs = append(runs, workspaceRuns...)
	}

	return runs, nil
}

func getWorkspaceRuns(client *tfe.Client, workspace *tfe.Workspace, runs map[string][]*tfe.Run) error {

	currentPage := 0
	totalPages := 1
	pageSize := 30

	for currentPage < totalPages {
		runPage, err := getRunPage(client, workspace.ID, tfe.RunListOptions{
			ListOptions: tfe.ListOptions{
				PageNumber: currentPage,
				PageSize:   pageSize,
			},
		})

		if err != nil { return err }

		for _, run := range runPage.Items {

			if run.Status == "applied" {
				key := run.StatusTimestamps.AppliedAt.AddDate(0,0, (run.StatusTimestamps.AppliedAt.Day() - 1) * -1).Format("2006-01-02")

				if list, ok := runs[key]; ok {
					runs[key] = append(list, run)
				}
			}

			//t := run.StatusTimestamps.AppliedAt.Round(time.)
			//
			//	print(t)

			//if run.Status == "applied" && run.CostEstimate != nil {



				//currentRun := &Record {
				//	OrganizationName:      workspace.Organization.Name,
				//	WorkspaceId:           workspace.ID,
				//	WorkspaceName:         workspace.Name,
				//	RunId:                 run.ID,
				//	CostEstimateId:        run.CostEstimate.ID,
				//	IsDestroy:			   run.IsDestroy,
				//	ProviderCostBreakdown: nil,
				//	StartTime:             run.StatusTimestamps.AppliedAt,
				//	EndTime:               time.Time{},
				//}
				//
				//if previousRun != nil {
				//	currentRun.EndTime = previousRun.StartTime
				//}
				//
				//// Destroy Runs are necessary for calculating Start/End times of Apply runs
				//// However, they are not needed in the final dataset. Filter out destroy records.
				//if !currentRun.IsDestroy {
				//	runs = append(runs, currentRun)
				//}
				//previousRun = currentRun
			//}
		}

		totalPages = runPage.TotalPages
		currentPage++
	}

	return nil
}

func getRunPage(client *tfe.Client, workspaceId string, options tfe.RunListOptions) (*tfe.RunList, error) {
	runs, err := client.Runs.List(context.Background(), workspaceId, options)
	if err != nil {
		return nil, err
	}

	return runs, nil
}
