package main

type Repository struct {
	Branches []*Branch
}

type LocalRepository struct {
	Name         string
	Head         *Commit
	Origin       *RemoteRepository
	ActiveBranch *Branch
	Repository   *Repository
}

type RemoteRepository struct {
	Name          string
	Url           string
	DefaultBranch *Branch
	Repository    *Repository
}

type Commit struct {
	Message string
	Date    string
}

type Branch struct {
	Name    string
	Commits []*Commit
}
