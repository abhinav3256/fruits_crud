package pkg

type Fruit struct {
	Id        int    `json:"id" `
	FruitName string `json:"fruit_name" binding:"required"`
}
