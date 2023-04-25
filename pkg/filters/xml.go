package filters

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	log "github.com/gothew/l-og"
)

type Adenda struct {
	XMLName xml.Name `xml:"http://www.sat.gob.gt/dte/fel/0.1.0 GTDocumento"`
	Campos  []Campo  `xml:"SAT>Adenda>ECFD>Personalizados>campoString"`
}

type Campo struct {
	XMLName xml.Name `xml:""`
	Nombre  string   `xml:"name,attr"`
	Valor   string   `xml:",chardata"`
}

func FilterNumeroInternoXml(fileName string) {
	//fileName := "dumps/test.xml"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	var adenda Adenda

	decoder, _ := ioutil.ReadAll(file)

	xml.Unmarshal(decoder, &adenda)

	for _, campo := range adenda.Campos {
		if campo.Nombre == "NUMEROINTERNO" {
			log.Infof("El número interno es: %s", campo.Valor)
		}
	}
}
