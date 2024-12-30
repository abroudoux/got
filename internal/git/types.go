package git

type Repository struct {
	Name string
	Head *Commit
	Branches []*Branch
	ActiveBranch *Branch
	Origin *RemoteRepository
}

type RemoteRepository struct {
	Name string
	Url string
	Repository *Repository
}

type Commit struct {
	Id string
	Message string
	Date string
}

type Branch struct {
	Name string
	Commits []*Commit
}

type repository interface {
	Init(repositoryName string)
	Commit(message string)
	LogCommits()
	Branch(branchName string)
	Checkout(branchName string)
	Merge(branchName string)
	RemoteAdd(remoteRepository *RemoteRepository)
}

type remoteRepository interface {
	CreateRemoteRepository(repositoryName string)
}