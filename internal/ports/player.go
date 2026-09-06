package ports

import "go.mod/internal/domain"

type PlayerService interface {
	InsertPlayer(player domain.Player) (id interface{}, err error)
}
