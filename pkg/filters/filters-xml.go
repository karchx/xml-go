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
	"os"

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
	}
}

func GetFileDate() string {
	filename := "../dumps/FACT_00a66b19-c583-4312-b739-67bbc7860225__a70dd235-5c1d-46ef-be02-9ba5e4a3b93c.xml"

	file, err := os.Stat(filename)

	if err != nil {
		log.Error(err)
	}
	return utils.GetDateFile(file)
}
