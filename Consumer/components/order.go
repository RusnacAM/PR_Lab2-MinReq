package components

type Order struct {
	Id       int     `json:"id"`
	Items    []int   `json:"items"`
	Priority int     `json:"priority"`
	MaxWait  float32 `json:"max_wait"`
}
