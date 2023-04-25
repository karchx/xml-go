// FILTER: NUMEROINTERNO
// FILTER: FECHA:HORA
// CVS:
// NUMEROINTERNO | FECHA
// XML:
//<ecfd:Personalizados>
//<ecfd:campoString name="_strartofdoc"></ INICIO ></ecfd:campoString>
//<ecfd:campoString name="_endfrases"></ FIN FRASES /></ecfd:campoString>
//<ecfd:campoString name="EXP"/>
//<ecfd:campoString name="SERIEINTERNA">XXX</ecfd:campoString>
//<ecfd:campoString name="NUMEROINTERNO">XXXXXX</ecfd:campoString>
//</ecfd:Personalizados>

package filters

import "github.com/karchx/xml-go/pkg/utils"

func GetFiltersCsv(csv string) []utils.CSV {
	return utils.ReadCsvFile(csv)
}
