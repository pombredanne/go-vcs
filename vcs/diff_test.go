	"sourcegraph.com/sourcegraph/go-vcs/vcs"
	t.Parallel()
			repo: makeGitRepositoryCmd(t, cmds...),
			repo: makeHgRepositoryCmd(t, hgCommands...),
	t.Parallel()
			baseRepo: makeGitRepositoryCmd(t, cmds...),
			headRepo: makeGitRepositoryCmd(t, cmds...),