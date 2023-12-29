package models

import (
	"encoding/xml"
	"time"
)

/*
type NFeXml struct {
	XMLName xml.Name `xml:"NfeProc"`
	NfeProc NfeProc  `xml:"nfeProc"`
	Versao  string   `xml:"versao"`
}
*/
type NfeProc struct {
	XMLName xml.Name `xml:"nfeProc"`
	Versao  string   `xml:"versao,attr"`
	Nfe     NFe      `xml:"NFe"`
	ProtNFe ProtNFe  `xml:"protNFe"`
}

type ProtNFe struct {
	XMLName xml.Name `xml:"protNFe"`
	InfProt InfProt  `xml:"infProt"`
}

type InfProt struct {
	XMLName xml.Name `xml:"infProt"`
	ChNFe   string   `xml:"chNFe"`
	CStat   int      `xml:"cStat"`
}

type NFe struct {
	XMLName xml.Name `xml:"NFe"`
	InfNFe  InfNFe   `xml:"infNFe"`
}

type InfNFe struct {
	XMLName xml.Name `xml:"infNFe"`
	Ide     Ide      `xml:"ide"`
	Emit    Emit     `xml:"emit"`
	Dest    Dest     `xml:"dest"`
	Det     []Det    `xml:"det"`
	Total   Total    `xml:"total"`
}

type Total struct {
	XMLName xml.Name `xml:"total"`
	ICMSTot ICMSTot  `xml:"ICMSTot"`
}

type ICMSTot struct {
	XMLName xml.Name `xml:"ICMSTot"`
	VNF     float64  `xml:"vNF"`
	VDesc   float64  `xml:"vDesc"`
	VSeg    float64  `xml:"vSeg"`
	VFrete  float64  `xml:"vFrete"`
	VProd   float64  `xml:"vProd"`
	VOutro  float64  `xml:"vOutro"`
	VCOFINS float64  `xml:"vCOFINS"`
	VPIS    float64  `xml:"vPIS"`
	VBC     float64  `xml:"vBC"`
	VIPI    float64  `xml:"vIPI"`
	VICMS   float64  `xml:"vICMS"`
	VBCST   float64  `xml:"vBCST"`
	VST     float64  `xml:"vST"`
}

type Ide struct {
	XMLName     xml.Name  `xml:"ide"`
	CUF         int       `xml:"cUF"`
	CNF         string    `xml:"cNF"`
	NatOp       string    `xml:"natOp"`
	Mod         string    `xml:"mod"`
	Serie       string    `xml:"serie"`
	NF          string    `xml:"nNF"`
	DhEmi       time.Time `xml:"dhEmi"`
	TpNF        int       `xml:"tpNF"`
	IdDest      int       `xml:"idDest"`
	CMunFG      int       `xml:"cMunFG"`
	TpImp       int       `xml:"tpImp"`
	TpEmis      int       `xml:"tpEmis"`
	CDV         int       `xml:"cDV"`
	TpAmb       int       `xml:"tpAmb"`
	FinNFe      int       `xml:"finNFe"`
	IndFinal    int       `xml:"indFinal"`
	IndPres     int       `xml:"indPres"`
	IndIntermed int       `xml:"indIntermed"`
	ProcEmi     int       `xml:"procEmi"`
	VerProc     string    `xml:"verProc"`
}

type Emit struct {
	XMLName xml.Name  `xml:"emit"`
	CNPJ    string    `xml:"CNPJ"`
	Nome    string    `xml:"xNome"`
	IE      string    `xml:"IE"`
	CRT     string    `xml:"CRT"`
	Ender   EnderEmit `xml:"enderEmit"`
}

type EnderEmit struct {
	XMLName xml.Name `xml:"enderEmit"`
	Lgr     string   `xml:"xLgr"`
	Nro     string   `xml:"nro"`
	Cpl     string   `xml:"xCpl"`
	Bairro  string   `xml:"xBairro"`
	CMun    int      `xml:"cMun"`
	Mun     string   `xml:"xMun"`
	UF      string   `xml:"UF"`
	CEP     string   `xml:"CEP"`
	CPais   int      `xml:"cPais"`
	Pais    string   `xml:"xPais"`
	Fone    string   `xml:"fone"`
}

type Dest struct {
	CNPJ      string    `xml:"CNPJ"`
	CPF       string    `xml:"CPF"`
	Nome      string    `xml:"xNome"`
	IE        string    `xml:"IE"`
	RG        string    `xml:"RG"`
	Email     string    `xml:"email"`
	IndIEDest int       `xml:"indIEDest"`
	Ender     EnderDest `xml:"enderDest"`
}

type EnderDest struct {
	XMLName xml.Name `xml:"enderDest"`
	Lgr     string   `xml:"xLgr"`
	Nro     string   `xml:"nro"`
	Cpl     string   `xml:"xCpl"`
	Bairro  string   `xml:"xBairro"`
	CMun    int      `xml:"cMun"`
	Mun     string   `xml:"xMun"`
	UF      string   `xml:"UF"`
	CEP     string   `xml:"CEP"`
	CPais   int      `xml:"cPais"`
	Pais    string   `xml:"xPais"`
	Fone    string   `xml:"fone"`
}
type Det struct {
	XMLName xml.Name `xml:"det"`
	Item    string   `xml:"nItem,attr"`
	Prod    Prod     `xml:"prod"`
}

type Prod struct {
	XMLName xml.Name `xml:"prod"`
	CProd   string   `xml:"cProd"`
	EAN     string   `xml:"cEAN"`
	Prod    string   `xml:"xProd"`
	NCM     string   `xml:"NCM"`
	CEST    string   `xml:"CEST"`
	CFOP    string   `xml:"CFOP"`
	UCom    string   `xml:"uCom"`
	QCom    float64  `xml:"qCom"`
	UnCom   float64  `xml:"vUnCom"`
	VProd   float64  `xml:"vProd"`
	EANTrib string   `xml:"cEANTrib"`
	UTrib   string   `xml:"uTrib"`
	QTrib   float64  `xml:"qTrib"`
	UnTrib  float64  `xml:"vUnTrib"`
	IndTot  int      `xml:"indTot"`
}
