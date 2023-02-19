package saq

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
	name              string
	product_link      string
	catagory          string
	saq_code          string
	volume            string
	country_of_origin string
	price             string
	rating_summary    string
	rating_actions    string
	bottled_in_quebec bool
}

type SearchResults struct {
	list  chan ProductInfo
	query string
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
