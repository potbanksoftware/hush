package archive

import "github.com/potbanksoftware/hush"

func (m *Manager) Uninstall(params hush.UninstallParams) error {
	// install folder is getting wiped anyway, nothing
	// in particular to do here.
	return nil
}
