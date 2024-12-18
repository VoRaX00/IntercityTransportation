package ticket

import "log/slog"

type Repo interface {
}

type Ticket struct {
	log  *slog.Logger
	repo Repo
}

func New(log *slog.Logger, repo Repo) *Ticket {
	return &Ticket{
		log:  log,
		repo: repo,
	}
}
