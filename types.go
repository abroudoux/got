package main

type Repository struct {
	Name         string
	Head         *Commit
	Branches     []*Branch
	ActiveBranch *Branch
	Origin       string
}

type RemoteRepository struct {
	Name          string
	Url           string
	DefaultBranch *Branch
	Repository    *Repository
}

type Commit struct {
	Id      string
	Message string
	Date    string
}

type Branch struct {
	Name       string
	Commits    []*Commit
	LastCommit *Commit
}
