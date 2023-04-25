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

func GetFiles(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatalf("%s", err)
	}
	for _, f := range files {
		log.Infof("FileName: %s", f.Name())
		log.Infof("FileDate: %s", utils.GetDateFile(f))
	}
}
