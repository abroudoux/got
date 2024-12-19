package git

type Repository struct {
	Name string
	Head *Commit
	Branches []*Branch
	ActiveBranch *Branch
}

type Commit struct {
	Id string
	Message string
	Date string
}

type Branch struct {
	Name string
	Commits []*Commit
	LastCommit *Commit
}