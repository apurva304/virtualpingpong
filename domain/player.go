package domain

type PlayerDetail struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	DefenceArrayLen int    `json:"defence_array_len"`
}
