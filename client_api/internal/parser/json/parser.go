package json

import (
	"encoding/json"
	"github.com/aelmel/ports-infra/client_api/internal/domain"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
)

func MonitorPath(basepath string) {
	for {
		time.Sleep(1 * time.Minute)
		log.Info("Checking %s", basepath)
		files, err := ioutil.ReadDir(basepath)
		if err != nil {
			log.Error("error reading dir %s error %v", basepath, err)
			os.Exit(1)
		}

		for _, f := range files {
			if filepath.Ext(f.Name()) == ".json" {
				readJsonFile(path.Join(basepath, f.Name()))
			}
		}
	}
}

func readJsonFile(fileName string) {
	file, _ := os.Open(fileName)
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Token()
	for decoder.More() {
		_, err := decoder.Token()
		if err != nil {
			log.Warn("error getting token", err)
		}
		var portDetails domain.PortDetails

		err = decoder.Decode(&portDetails)
		if err != nil {
			log.Warn("error parsing line %v", err)
		}
	}
	decoder.Token()
	os.Rename(fileName, fileName+".done")
}
