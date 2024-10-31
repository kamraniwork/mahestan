package serializers

type InputCategorySerializer struct {
	Name        string `json:"name" validate:"required,max=20,gt=3"`
	Description string `json:"description" validate:"required,max=255"`
}

type OutputDetailCategorySerializer struct {
	Name        string `json:"name" validate:"required,max=20,gt=3"`
	Description string `json:"description" validate:"required,max=255"`
}

type OutputListCategorySerializer struct {
	Category []OutputDetailCategorySerializer `json:"category"`
}
