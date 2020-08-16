package datalib

// Аппаратное обеспечение (Аппаратный состав)
type Hardware struct {
	id string
	string
}

type cpu struct {
	Manufacturer string `xml:"manufacturer"`
	Name         string `xml:"name"`
	Speed        uint32 `xml:"speed"`
	Cores        uint32 `xml:"cores"`
	Thread       uint32 `xml:"thread"`
}

type Software struct {
	id string
}

type Perfomance struct {
	id string
}

type ErrorList struct {
	id string
}

type PCSpec struct {
	id string
}
