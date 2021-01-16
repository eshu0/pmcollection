package pmcollection

import (
	sl "github.com/eshu0/simplelogger/pkg"
)

type PostmanManager struct {
	sl.AppLogger
	Collections []*Collection
}

func (pm *PostmanManager) SaveColletion(filepath string) {

	for _, col := range pm.Collections {
		success := col.Save(filepath, pm.Log)

		if success {
			pm.LogInfo("SaveColletion", "Saved "+col.Info.Name)
		}
	}
}

func (pm *PostmanManager) LoadColletion(filepath string) {

	pmcol := NewCollection("somename")
	cols := pm.Collections

	pm.LogDebug("LoadColletion", "Started")
	col, success := pmcol.Load(filepath, pm.Log)

	if success {
		pm.LogInfo("LoadColletion", "Loaded "+col.Info.Name)
		cols = append(cols, col)
	}
	pm.Collections = cols
}

func (pm *PostmanManager) PrintCollections() {

	pm.LogInfo("PrintCollections", "Started printing collections")

	for _, col := range pm.Collections {

		if col.Info != nil {
			pm.LogInfo("PrintCollections", col.Info.Name)
			pm.LogInfo("PrintCollections", col.Info.Postman_Id)
			pm.LogInfo("PrintCollections", col.Info.Schema)
		}

		for _, item := range col.Item {
			pm.LogInfo("PrintCollections", item.Name)
		}
	}
}
