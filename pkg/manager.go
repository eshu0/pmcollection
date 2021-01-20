package pmcollection

import (
	"fmt"

	sl "github.com/eshu0/simplelogger/pkg"
)

type PostmanManager struct {
	sl.AppLogger
	Collections map[string]*Collection
	//Folderpath  string
}

func NewPostmanManager() *PostmanManager {
	pm := PostmanManager{}
	collections := make(map[string]*Collection)
	pm.Collections = collections
	return &pm
}

func (pm *PostmanManager) SaveColletions(folderpath string) {
	for k, _ := range pm.Collections {
		pm.SaveColletion(fmt.Sprintf("%s%s.postman_collection.json", folderpath, pm.Collections[k].Info.Name))
	}

}

func (pm *PostmanManager) SaveColletion(filepath string) {

	for _, col := range pm.Collections {
		success := col.Save(filepath, pm.Log)

		if success {
			pm.LogInfo("SaveColletion", "Saved "+col.Info.Name)
		}
	}
}

func (pm *PostmanManager) AddItem(collectionname string, item *Item) {

	if pm.Collections[collectionname] == nil {
		pm.LogDebug("AddItem", "Adding "+collectionname)
		col := NewCollection(collectionname)
		pm.AddColletion(col)
	}

	items := pm.Collections[collectionname].Item
	items = append(items, item)
	pm.Collections[collectionname].Item = items
}

func (pm *PostmanManager) AddColletion(col *Collection) {

	if pm.Collections[col.Info.Name] != nil {
		pm.LogDebug("AddColletion", "Updating "+col.Info.Name)
		pm.Collections[col.Info.Name] = col

	} else {
		/*
			cols := pm.Collections

			pm.LogDebug("AddColletion", "Adding "+col.Info.Name)
			cols = append(cols, col)
			pm.Collections = cols
		*/
		pm.LogDebug("AddColletion", "Adding "+col.Info.Name)

		pm.Collections[col.Info.Name] = col
	}

}

func (pm *PostmanManager) LoadColletion(filepath string) {

	pmcol := NewCollection("somename")
	pm.LogDebug("LoadColletion", "Started")
	col, success := pmcol.Load(filepath, pm.Log)

	if success {
		pm.LogInfo("LoadColletion", "Loaded "+col.Info.Name)
		pm.AddColletion(col)
	}
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
