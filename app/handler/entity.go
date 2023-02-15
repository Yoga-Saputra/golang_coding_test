package handler

type LanguangeEntity struct {
	Appeared       int      `json:"appeared"`
	Languages      string   `json:"language"`
	Created        []string `json:"created"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation"`
}

type GetLanguangeDetailInput struct {
	ID *int `uri:"id" binding:"required"`
}

type Relation struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences   []string `json:"influences"`
}

type LanguagesOnlyEntity struct {
	Languages string `json:"language"`
}

type Data1 struct {
	Input string
}

type Data2 struct {
	Input string
}
