package json

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aelmel/ports-infra/client_api/internal/client/port"
	"github.com/aelmel/ports-infra/client_api/internal/domain"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
)

type Parser interface {
	Monitor()
}

type parser struct {
	basepath string
	client   port.Client
}

func NewParser(client port.Client, basepath string) Parser {
	return parser{basepath: basepath, client: client}
}

func (p parser) Monitor() {
	for {
		time.Sleep(10 * time.Second)
		log.Info("Checking ", p.basepath)
		files, err := ioutil.ReadDir(p.basepath)
		if err != nil {
			log.Error("error reading dir %s error %v", p.basepath, err)
			os.Exit(1)
		}

		for _, f := range files {
			if filepath.Ext(f.Name()) == ".json" {
				p.readJsonFile(path.Join(p.basepath, f.Name()))
			}
		}
	}
}

func (p parser) readJsonFile(fileName string) {
	log.Info(fmt.Sprintf("Processing file %s", fileName))
	file, err := os.Open(fileName)
	if err != nil {
		log.Info("error opening file ", fileName)
	}
	defer file.Close()
	defer os.Rename(fileName, fileName+".done")

	decoder := json.NewDecoder(file)
	decoder.Token()
	for decoder.More() {
		key, err := decoder.Token()
		_ = key
		if err != nil {
			log.Warn("error getting token", err)
			continue
		}
		var portDetails domain.PortDetails

		err = decoder.Decode(&portDetails)
		if err != nil {
			log.Warn("error parsing line %v", err)
			continue
		}
		err = p.client.AddPort(context.Background(), key.(string), portDetails)
		if err != nil {
			log.Warn(fmt.Sprintf("error add port %v error %v", portDetails, err))
			continue
		}
	}
	decoder.Token()
}
