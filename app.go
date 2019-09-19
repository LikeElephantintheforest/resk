package xffresk

import (
	"xffresk/infra"
	"xffresk/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
}