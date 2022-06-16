# TFE Usage Stats

```
./tfe-usage-stats -host="https://app.terraform.io" -token="" --gen-workspace-owner-sheet=true --organization=""
```
# Optional Flags

Generate Workspace Owners Sheet will generate a spreadsheet in the current directory that lists out all the team/workspace relationships in the respective organizations

```
--gen-workspace-owner-sheet=true 
```

The Organization flag will limit the scope to a specific organization. This can help with really large deployments of TFC/E

```
--organization="tfc-peyton"
```

# Example

<p align="center">
    <img align="center" src="images/example.png" alt="example"/>
</p>

# Build

```
go build cmd/tfe-usage-stats.go
```
