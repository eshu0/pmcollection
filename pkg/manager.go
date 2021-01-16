package pmcollection

import (
	sl "github.com/eshu0/simplelogger/pkg"
)

type PostmanManager struct {
	sl.AppLogger
	Collections []*Collection
}

func (pm *PostmanManager) Load(filepath string) {

	pmcol := NewCollection("todo")
	return pmcol.Load(filepath, pm.Logger)

}
