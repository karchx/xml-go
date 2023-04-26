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

	"github.com/karchx/xml-go/pkg/utils"
)

func GetFiles(directory string, filters []utils.CSV) {

	var index = 1
	files, _ := ioutil.ReadDir(directory)
	for i := 1; i < len(filters); i++ {
		for _, f := range files {

			fileName := directory + "/" + f.Name()

			FilterNumeroInternoXml(fileName, filters[index].NumeroInterno)
		}
		index++
	}
}
