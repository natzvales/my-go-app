package books

type CreateBookDTO struct {
	Title  string `json:"title" binding:"required" message:"Title is required"`
	Author string `json:"author" binding:"required" message:"Author is required"`
	// ISBN   string `json:"isbn" binding:"required"`
}

type UpdateBookDTO struct {
	Title  string `json:"title" binding:"required" message:"Title is required"`
	Author string `json:"author" binding:"required" message:"Author is required"`
	// ISBN   string `json:"isbn" binding:"required"`
}
