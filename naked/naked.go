package naked

import "github.com/potbanksoftware/hush"

type Manager struct {
}

var _ hush.Manager = (*Manager)(nil)

func (m *Manager) Name() string {
	return "naked"
}
