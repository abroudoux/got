package git

func ReverseCommitsOrder(commits []*Commit) ([]*Commit, error) {
    var reversedCommits []*Commit
    for i := len(commits) - 1; i >= 0; i-- {
        reversedCommits = append(reversedCommits, commits[i])
    }
    return reversedCommits, nil
}

func GetLastCommit(branch *Branch) *Commit {
    return branch.Commits[len(branch.Commits) - 1]
}