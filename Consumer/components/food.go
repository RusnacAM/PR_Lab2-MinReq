package components

type Food struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	PreparationTime  int    `json:"preparationTime"`
	Complexity       int    `json:"complexity"`
	CookingApparatus string `json:"cookingApparatus"`
}
