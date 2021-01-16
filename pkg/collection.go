package pmcollection

import (
	"encoding/json"
	"io/ioutil"
	"os"

	slinterfaces "github.com/eshu0/simplelogger/pkg/interfaces"
)

//Info information about the colelction
type Info struct {
	Name string `json:"name"`

	Postman_Id string `json:"_postman_id"`

	//"https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	Schema string `json:"schema"`
}

//Collection top level
type Collection struct {
	Info *Info `json:"info"`

	Item []*Item `json:"item"`
}

func NewCollection(name string) *Collection {

	info := &Info{Name: name, Postman_Id: "0f77ef88-ffff-4359-9e56-b19530b7378b", Schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"}
	col := Collection{Info: info}

	return &col
}

func (rsc *Collection) Save(ConfigFilePath string, Log slinterfaces.ISimpleLogger) bool {
	bytes, err1 := json.MarshalIndent(rsc, "", "\t") //json.Marshal(p)
	if err1 != nil {
		Log.LogErrorf("SaveToFile()", "Marshal json for %s failed with %s ", ConfigFilePath, err1.Error())
		return false
	}

	err2 := ioutil.WriteFile(ConfigFilePath, bytes, 0644)
	if err2 != nil {
		Log.LogErrorf("SaveToFile()", "Saving %s failed with %s ", ConfigFilePath, err2.Error())
		return false
	}

	return true

}

func (rsc *Collection) Load(ConfigFilePath string, Log slinterfaces.ISimpleLogger) (*Collection, bool) {
	ok, err := rsc.checkFileExists(ConfigFilePath)
	if ok {
		bytes, err1 := ioutil.ReadFile(ConfigFilePath) //ReadAll(jsonFile)
		if err1 != nil {
			Log.LogErrorf("LoadFile()", "Reading '%s' failed with %s ", ConfigFilePath, err1.Error())
			return nil, false
		}

		rserverconfig := &Collection{}

		err2 := json.Unmarshal(bytes, rserverconfig)

		if err2 != nil {
			Log.LogErrorf("LoadFile()", " Loading %s failed with %s ", ConfigFilePath, err2.Error())
			return nil, false
		}

		if rserverconfig != nil && rserverconfig.Info != nil {
			Log.LogDebugf("LoadFile()", "Read %v ", rserverconfig)

			Log.LogDebugf("LoadFile()", "Read %s ", rserverconfig.Info.Name)
		} else {
			Log.LogDebugf("LoadFile()", "Read %v ", rserverconfig)
		}

		//rs.Log.LogDebugf("LoadFile()", "Port in config %s ", rs.Config.Port)

		return rserverconfig, true
	} else {

		if err != nil {
			Log.LogErrorf("LoadFile()", "'%s' was not found to load with error: %s", ConfigFilePath, err.Error())
		} else {
			Log.LogErrorf("LoadFile()", "'%s' was not found to load", ConfigFilePath)
		}

		return nil, false
	}
}

func (rsc *Collection) checkFileExists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, err
	}
	return !info.IsDir(), nil
}
