// FILTER: NUMEROINTERNO
// XML:
//<ecfd:Personalizados>
//<ecfd:campoString name="_strartofdoc"></ INICIO ></ecfd:campoString>
//<ecfd:campoString name="_endfrases"></ FIN FRASES /></ecfd:campoString>
//<ecfd:campoString name="EXP"/>
//<ecfd:campoString name="SERIEINTERNA">XXX</ecfd:campoString>
//<ecfd:campoString name="NUMEROINTERNO">XXXXXX</ecfd:campoString>
//</ecfd:Personalizados>
package filters

import (
	"io/ioutil"

	log "github.com/gothew/l-og"
	"github.com/karchx/xml-go/pkg/utils"
)

func GetFiles(directory string, filters []utils.CSV) {
	files, err := ioutil.ReadDir(directory)
	var index int

	if err != nil {
		log.Fatalf("%s", err)
	}
	for i, f := range files {
		fileName := directory + "/" + f.Name()

		if i >= len(filters)-1 {
			i = 0
		}
		index = i + 1

		FilterNumeroInternoXml(fileName, filters[index].NumeroInterno)
	}
}
