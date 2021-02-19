package auction

// Auction ...
type Auction struct {
	Hash         string             `json:"hash"`
	MinMaxValues map[string]float64 `json:"minMaxValues"`
	Server       []Server           `json:"server"`
}

// Server ...
type Server struct {
	Key                 int64        `json:"key"`
	Name                string       `json:"name"`
	Description         []string     `json:"description"`
	CPU                 string       `json:"cpu"`
	CPUBenchmark        int64        `json:"cpu_benchmark"`
	CPUCount            int64        `json:"cpu_count"`
	IsHighio            bool         `json:"is_highio"`
	IsECC               bool         `json:"is_ecc"`
	Traffic             Traffic      `json:"traffic"`
	Dist                []Dist       `json:"dist"`
	Bandwith            int64        `json:"bandwith"`
	RAM                 int64        `json:"ram"`
	Price               string       `json:"price"`
	PriceV              string       `json:"price_v"`
	RAMHr               RAMHr        `json:"ram_hr"`
	SetupPrice          string       `json:"setup_price"`
	HDDSize             int64        `json:"hdd_size"`
	HDDCount            int64        `json:"hdd_count"`
	HDDHr               string       `json:"hdd_hr"`
	FixedPrice          bool         `json:"fixed_price"`
	NextReduce          int64        `json:"next_reduce"`
	NextReduceHr        string       `json:"next_reduce_hr"`
	NextReduceTimestamp int64        `json:"next_reduce_timestamp"`
	Datacenter          []Datacenter `json:"datacenter"`
	Specials            []Special    `json:"specials"`
	SpecialHDD          SpecialHDD   `json:"specialHdd"`
	Freetext            string       `json:"freetext"`
}

// Datacenter ...
type Datacenter string

// foo
const (
	Fsn      Datacenter = "FSN"
	Fsn1Dc1  Datacenter = "FSN1-DC1"
	Fsn1Dc10 Datacenter = "FSN1-DC10"
	Fsn1Dc11 Datacenter = "FSN1-DC11"
	Fsn1Dc12 Datacenter = "FSN1-DC12"
	Fsn1Dc13 Datacenter = "FSN1-DC13"
	Fsn1Dc14 Datacenter = "FSN1-DC14"
	Fsn1Dc15 Datacenter = "FSN1-DC15"
	Fsn1Dc16 Datacenter = "FSN1-DC16"
	Fsn1Dc3  Datacenter = "FSN1-DC3"
	Fsn1Dc4  Datacenter = "FSN1-DC4"
	Fsn1Dc5  Datacenter = "FSN1-DC5"
	Fsn1Dc6  Datacenter = "FSN1-DC6"
	Fsn1Dc7  Datacenter = "FSN1-DC7"
	Fsn1Dc8  Datacenter = "FSN1-DC8"
	Hel      Datacenter = "HEL"
	Hel1Dc2  Datacenter = "HEL1-DC2"
	Hel1Dc4  Datacenter = "HEL1-DC4"
	Nbg      Datacenter = "NBG"
	Nbg1Dc1  Datacenter = "NBG1-DC1"
)

// Dist ...
type Dist string

// foo
const (
	RescueSystem Dist = "Rescue system"
)

// RAMHr ...
type RAMHr string

// foo
const (
	The128GB RAMHr = "128 GB"
	The12GB  RAMHr = "12 GB"
	The160GB RAMHr = "160 GB"
	The16GB  RAMHr = "16 GB"
	The192GB RAMHr = "192 GB"
	The24GB  RAMHr = "24 GB"
	The256GB RAMHr = "256 GB"
	The2GB   RAMHr = "2 GB"
	The32GB  RAMHr = "32 GB"
	The384GB RAMHr = "384 GB"
	The48GB  RAMHr = "48 GB"
	The4GB   RAMHr = "4 GB"
	The512GB RAMHr = "512 GB"
	The64GB  RAMHr = "64 GB"
	The8GB   RAMHr = "8 GB"
)

// SpecialHDD ...
type SpecialHDD string

// foo
const (
	DatacenterSSD           SpecialHDD = "Datacenter SSD"
	Empty                   SpecialHDD = ""
	SpecialHDDEntHDD        SpecialHDD = "Ent. HDD"
	SpecialHDDNVMeSSD       SpecialHDD = "NVMe SSD"
	The1XSSD120GB           SpecialHDD = "+ 1x SSD  120 GB"
	The1XSSD240GB           SpecialHDD = "+ 1x SSD  240 GB"
	The1XSSD240GBDatacenter SpecialHDD = "+ 1x SSD  240 GB Datacenter"
	The1XSSD250GB           SpecialHDD = "+ 1x SSD  250 GB"
	The1XSSD256GB           SpecialHDD = "+ 1x SSD  256 GB"
	The1XSSD480GBDatacenter SpecialHDD = "+ 1x SSD  480 GB Datacenter"
	The1XSSD500GB           SpecialHDD = "+ 1x SSD  500 GB"
	The1XSSD512GB           SpecialHDD = "+ 1x SSD  512 GB"
	The1XSSD960GBDatacenter SpecialHDD = "+ 1x SSD  960 GB Datacenter"
	The1XSSDM2NVMe512GB     SpecialHDD = "+ 1x SSD M.2 NVMe 512 GB"
	The2XSSD120GB           SpecialHDD = "+ 2x SSD  120 GB"
	The2XSSD240GB           SpecialHDD = "+ 2x SSD  240 GB"
	The2XSSD240GBDatacenter SpecialHDD = "+ 2x SSD  240 GB Datacenter"
	The2XSSD250GB           SpecialHDD = "+ 2x SSD  250 GB"
	The2XSSD256GB           SpecialHDD = "+ 2x SSD  256 GB"
	The2XSSD480GBDatacenter SpecialHDD = "+ 2x SSD  480 GB Datacenter"
	The2XSSD500GB           SpecialHDD = "+ 2x SSD  500 GB"
	The2XSSD512GB           SpecialHDD = "+ 2x SSD  512 GB"
	The2XSSD960GBDatacenter SpecialHDD = "+ 2x SSD  960 GB Datacenter"
	The2XSSDM2NVMe512GB     SpecialHDD = "+ 2x SSD M.2 NVMe 512 GB"
)

// Special ...
type Special string

// foo ...
const (
	DcSSD          Special = "DC SSD"
	ECC            Special = "ECC"
	Hwr            Special = "HWR"
	INIC           Special = "iNIC"
	RedPS          Special = "Red.PS"
	SAS            Special = "SAS"
	SSD            Special = "SSD"
	SpecialEntHDD  Special = "Ent. HDD"
	SpecialNVMeSSD Special = "NVMe SSD"
)

// Traffic ...
type Traffic string

// foo ...
const (
	Unlimited Traffic = "unlimited"
)
