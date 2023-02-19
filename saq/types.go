package saq

import "fmt"

type Language int8

const (
	Français Language = iota
	English
)

func (l Language) String() string {

	switch l {
	case Français:
		return "fr"
	case English:
		return "en"
	default:
		return "fr"
	}

}

type ProductInfo struct {
	Name            string
	ProductLink     string
	Catagory        string
	SaqCode         string
	Volume          string
	CountryOfOrigin string
	Price           string
	RatingSummary   string
	RatingActions   string
	BottledInQuebec bool
}

func (p ProductInfo) ToStringArray() []string {
	return []string{
		p.Name,
		p.Catagory,
		p.SaqCode,
		p.Volume,
		p.CountryOfOrigin,
		p.Price,
		p.RatingSummary,
		p.RatingActions,
		fmt.Sprintf("%v", p.BottledInQuebec),
		p.ProductLink,
	}

}

type SearchResults struct {
	List  chan ProductInfo
	Query string
}

type Filter struct {
	// catagory []Catagory;
	// region []Region;
	// price_minimum int32;
	// price_maximum int32;
	// availability []Availability;
	// product_of_quebec []QuebecProduct;
	// taste_tag []TasteTag;
	// customer_rating int32; // 1 2 3 4 5
	// type_of_spirit: []TypeOfSpirit;
	// min_sugar_content: int32;
	// max_sugar_content: int32;
	// min_degree_of_alcohol: int32;
	// max_degree_of_alcohol: int32;
	// special_feature: []SpecialFeature;
	// size: []Size;
	// designation_of_origin: []DesignationOfOrigin;
	// grape_variety: []GrapeVariety;
	// vintage: []int32;
}
