package installers

import (
	"github.com/potbanksoftware/hush"
	"github.com/potbanksoftware/hush/archive"
	"github.com/potbanksoftware/hush/naked"
)

func GetManager(typ hush.InstallerType) hush.Manager {
	switch typ {
	case hush.InstallerTypeArchive:
		return &archive.Manager{}
	case hush.InstallerTypeNaked:
		return &naked.Manager{}
	default:
		return nil
	}
}
