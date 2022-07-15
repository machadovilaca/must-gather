package main

import "github.com/go-git/go-git/v5"

const (
	name = "must-gather"

	gitOwner     = "machadovilaca"
	gitTargetDir = "/tmp/release/" + gitOwner + "/" + name
	quayOwner    = "kubevirt"
)

type release struct {
	version     string
	gitHubToken string

	repository *git.Repository
}
