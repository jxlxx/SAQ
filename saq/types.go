package saq

type WineInfo struct {
	name              string
	productLink       string
	wine_type         string
	saq_code          string
	volume            string
	country_of_origin string
	price             float32
	customer_rating   int32
	bottled_in_quebec bool
}

type SearchResults struct {
	wines []WineInfo
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
