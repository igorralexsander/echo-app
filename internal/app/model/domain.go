package model

type DomainMapping struct {
	TargetSystem            string `json:"targetSystem"`
	OriginSystem            string `json:"originSystem"`
	OriginSystemDomainKey   string `json:"originSystemDomainKey"`
	OriginSystemDomainValue string `json:"originSystemDomainValue"`
	TargetSystemDomainValue string `json:"targetSystemDomainValue"`
}
