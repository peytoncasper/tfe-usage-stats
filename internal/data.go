package internal

import (
	tfe "github.com/hashicorp/go-tfe"
)

type Dataset struct {
	Groups map[string][]*tfe.Run
}

func NewDataset() Dataset {
	return Dataset{
		Groups: map[string][]*tfe.Run{},
	}
}