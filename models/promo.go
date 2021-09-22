package models

type Promo struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	Name      string  `json:"name"`
	Date_from string  `json:"date_from"`
	Date_to   string  `json:"date_to"`
	Available int     `json:"available"`
	Amount    float32 `json:"amount"`
	Allocated int     `json:"quantity_alocated"`
}

type CreatePromo struct {
	Name      string  `json:"name" binding:"required"`
	Date_from string  `json:"date_from" binding:"required"`
	Date_to   string  `json:"date_to" binding:"required"`
	Available int     `json:"available" binding:"required"`
	Amount    float32 `json:"amount" binding:"required"`
	Allocated int     `json:"alocated" binding:"required"`
}

type UpdatePromo struct {
	Name      string  `json:"name"`
	Date_from string  `json:"date_from"`
	Date_to   string  `json:"date_to"`
	Available int     `json:"available"`
	Amount    float32 `json:"amount"`
	Allocated int     `json:"alocated"`
}

type UsePromo struct {
	Available int `json:"available"`
}
