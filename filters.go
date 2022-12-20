package nostrate

import "time"

type Filters struct {
	Ids     []string  `json:"ids"`
	Authors []string  `json:"authors"`
	Kinds   []Kind    `json:"kinds"`
	Events  []string  `json:"#e"`
	PubKeys []string  `json:"#p"`
	Since   time.Time `json:"since"`
	Until   time.Time `json:"until"`
	Limit   int       `json:"limit"`
}
