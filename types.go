package main

type Repository struct {
	Name     string
	Branches []*Branch
}

type LocalRepository struct {
	Head         *Commit
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
	Id      string
	Message string
	Date    string
}

type Branch struct {
	Name       string
	Commits    []*Commit
	LastCommit *Commit
}
