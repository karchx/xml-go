package filters

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	log "github.com/gothew/l-og"
	"github.com/karchx/xml-go/pkg/fs"
)

type Adenda struct {
	XMLName xml.Name `xml:"GTDocumento"`
	Campos  []Campo  `xml:"SAT>Adenda>ECFD>Personalizados>campoString"`
}

type Campo struct {
	XMLName xml.Name `xml:""`
	Nombre  string   `xml:"name,attr"`
	Valor   string   `xml:",chardata"`
}

func FilterNumeroInternoXml(fileName string, key string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable open file: %v", err)
	}
	defer file.Close()

	var adenda Adenda

	decoder, _ := ioutil.ReadAll(file)

	xml.Unmarshal(decoder, &adenda)

	for _, campo := range adenda.Campos {
		if campo.Nombre == "NUMEROINTERNO" {
			if key == campo.Valor {
				log.Infof("NUMEROINTERNO: %s", campo.Valor)
				log.Infof("MATCH WITH FILE: %s", fileName)
				fs.MoveMatchNewFolder(fileName)
			}
		}
	}
}
