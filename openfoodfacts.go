package openfoodfacts

import "encoding/json"

type GetProductResponse struct {
	Code          string  `json:"code"`
	Status        int     `json:"status"`
	StatusVerbose string  `json:"status_verbose"`
	Product       Product `json:"product"`
}

type Product struct {
	ID                                          string         `json:"id"`
	Code                                        string         `json:"code"`
	Brands                                      string         `json:"brands"`
	BrandsTags                                  []string       `json:"brands_tags"`
	GenericName                                 string         `json:"generic_name"`
	GenericNameEn                               string         `json:"generic_name_en"`
	CreatedTime                                 int64          `json:"created_t"`
	LastImageTime                               int64          `json:"last_image_t"`
	LastModifiedTime                            int64          `json:"last_modified_t"`
	CompletedTime                               int64          `json:"completed_t"`
	ImageFrontSmallURL                          string         `json:"image_front_smallURL"`
	ImageFrontThumbURL                          string         `json:"image_front_thumb_url"`
	ImageFrontURL                               string         `json:"image_front_url"`
	ImageIngredientsSmallURL                    string         `json:"image_ingredients_small_url"`
	ImageIngredientsThumbURL                    string         `json:"image_ingredients_thumb_url"`
	ImageIngredientsURL                         string         `json:"image_ingredients_url"`
	ImageNutritionSmallURL                      string         `json:"image_nutrition_small_url"`
	ImageNutritionThumbURL                      string         `json:"image_nutrition_thumb_url"`
	ImageNutritionURL                           string         `json:"image_nutrition_url"`
	ImageSmallURL                               string         `json:"image_small_url"`
	ImageThumbURL                               string         `json:"image_thumb_url"`
	ImageURL                                    string         `json:"image_url"`
	Creator                                     string         `json:"creator"`
	Editors                                     []string       `json:"editors"`
	EditorsTags                                 []string       `json:"editors_tags"`
	Correctors                                  []string       `json:"correctors"`
	CorrectorsTags                              []string       `json:"correctors_tags"`
	Informers                                   []string       `json:"informers"`
	InformersTags                               []string       `json:"informers_tags"`
	Photographers                               []string       `json:"photographers"`
	PhotographersTags                           []string       `json:"photographers_tags"`
	LastEditDatesTags                           []string       `json:"last_edit_dates_tags"`
	LastEditor                                  string         `json:"last_editor"`
	LastImageDatesTags                          []string       `json:"last_image_dates_tags"`
	LastModifiedBy                              string         `json:"last_modified_by"`
	Additives                                   string         `json:"additives"`
	AdditivesDebugTags                          []string       `json:"additives_debug_tags"`
	AdditivesNumber                             int            `json:"additives_n"`
	AdditivesOldNumber                          int            `json:"additives_old_n"`
	AdditivesOldTags                            []string       `json:"additives_old_tags"`
	AdditivesPrev                               string         `json:"additives_prev"`
	AdditivesPrevNumber                         int            `json:"additives_prev_n"`
	AdditivesPrevTags                           []string       `json:"additives_prev_tags"`
	AdditivesTags                               []string       `json:"additives_tags"`
	Allergens                                   string         `json:"allergens"`
	AllergensHierarchy                          []string       `json:"allergens_hierarchy"`
	AllergensTags                               []string       `json:"allergens_tags"`
	Categories                                  string         `json:"categories"`
	CategoriesDebugTags                         []string       `json:"categories_debug_tags"`
	CategoriesHierarchy                         []string       `json:"categories_hierarchy"`
	CategoriesPrevHierarchy                     []string       `json:"categories_prev_hierarchy"`
	CategoriesPrevTags                          []string       `json:"categories_prev_tags"`
	CategoriesTags                              []string       `json:"categories_tags"`
	CitiesTags                                  []string       `json:"cities_tags"`
	CodesTags                                   []string       `json:"codes_tags"`
	Complete                                    int            `json:"complete"`
	Countries                                   string         `json:"countries"`
	CountriesHierarchy                          []string       `json:"countries_hierarchy"`
	CountriesTags                               []string       `json:"countries_tags"`
	EmbCodes                                    string         `json:"emb_codes"`
	EmbCodes20141016                            string         `json:"emb_codes_20141016"`
	EmbCodesOrig                                string         `json:"emb_codes_orig"`
	EmbCodesTags                                []string       `json:"emb_codes_tags"`
	EntryDatesTags                              []string       `json:"entry_dates_tags"`
	ExpirationDate                              string         `json:"expiration_date"`
	FruitsVegetablesNuts100GEstimate            json.Number    `json:"fruits-vegetables-nuts_100g_estimate"`
	Ingredients                                 []Ingredient   `json:"ingredients"`
	IngredientsDebug                            []string       `json:"ingredients_debug"` // TODO see if it works with null. id: 3660603004118
	IngredientsFromOrThatMayBeFromPalmOilNumber int            `json:"ingredients_from_or_that_may_be_from_palm_oil_n"`
	IngredientsFromPalmOilNumber                int            `json:"ingredients_from_palm_oil_n"`
	IngredientsIDsDebug                         []string       `json:"ingredients_ids_debug"`
	IngredientsN                                string         `json:"ingredients_n"`
	IngredientsNTags                            []string       `json:"ingredients_n_tags"`
	IngredientsTags                             []string       `json:"ingredients_tags"`
	IngredientsText                             string         `json:"ingredients_text"`
	IngredientsTextDebug                        string         `json:"ingredients_text_debug"`
	IngredientsTextEn                           string         `json:"ingredients_text_en"`
	IngredientsTextWithAllergens                string         `json:"ingredients_text_with_allergens"`
	IngredientsTextWithAllergensEn              string         `json:"ingredients_text_with_allergens_en"`
	IngredientsThatMayBeFromPalmOilNumber       int            `json:"ingredients_that_may_be_from_palm_oil_n"`
	IngredientsThatMayBeFromPalmOilTags         []string       `json:"ingredients_that_may_be_from_palm_oil_tags"`
	InterfaceVersionCreated                     string         `json:"interface_version_created"`
	InterfaceVersionModified                    string         `json:"interface_version_modified"`
	Keywords                                    []string       `json:"_keywords"`
	Labels                                      string         `json:"labels"`
	LabelsDebugTags                             []string       `json:"labels_debug_tags"`
	LabelsHierarchy                             []string       `json:"labels_hierarchy"`
	LabelsPrevHierarchy                         []string       `json:"labels_prev_hierarchy"`
	LabelsPrevTags                              []string       `json:"labels_prev_tags"`
	LabelsTags                                  []string       `json:"labels_tags"`
	Lang                                        string         `json:"lang"`
	Languages                                   map[string]int `json:"languages"`
	LanguagesCodes                              map[string]int `json:"languages_codes"`
	LanguagesHierarchy                          []string       `json:"languages_hierarchy"`
	LanguagesTags                               []string       `json:"languages_tags"`
	Locale                                      string         `json:"lc"`
	MaxImgID                                    string         `json:"max_imgid"`
	NewAdditivesNumber                          int            `json:"new_additives_n"`
	NoNutritionData                             string         `json:"no_nutrition_data"`
	NutrientLevels                              NutrientLevels `json:"nutrient_levels"`
	NutrientLevelsTags                          []string       `json:"nutrient_levels_tags"`
	Nutriments                                  Nutriments     `json:"nutriments"`
	NutritionDataPer                            string         `json:"nutrition_data_per"`
	NutritionGradeFr                            string         `json:"nutrition_grade_fr"`
	NutritionGrades                             string         `json:"nutrition_grades"`
	NutritionGradesTags                         []string       `json:"nutrition_grades_tags"`
	NutritionScoreDebug                         string         `json:"nutrition_score_debug"`
	OriginTags                                  []string       `json:"origins_tags"`
	Origins                                     string         `json:"origins"`
	Packaging                                   string         `json:"packaging"`
	PackagingTags                               []string       `json:"packaging_tags"`
	PnnsGroups1                                 string         `json:"pnns_groups_1"`
	PnnsGroups1Tags                             []string       `json:"pnns_groups_1_tags"`
	PnnsGroups2                                 string         `json:"pnns_groups_2"`
	PnnsGroups2Tags                             []string       `json:"pnns_groups_2_tags"`
	ProductName                                 string         `json:"product_name"`
	ProductNameEn                               string         `json:"product_name_en"`
	PurchasePlaces                              string         `json:"purchase_places"`
	PurchasePlacesTags                          []string       `json:"purchase_places_tags"`
	Quantity                                    string         `json:"quantity"`
	Rev                                         int            `json:"rev"`
	ScansNumber                                 int            `json:"scans_n"`
	ServingQuantity                             json.Number    `json:"serving_quantity"`
	ServingSize                                 string         `json:"serving_size"`
	SortKey                                     int            `json:"sortkey"`
	States                                      string         `json:"states"`
	StatesHierarchy                             []string       `json:"states_hierarchy"`
	StatesTags                                  []string       `json:"states_tags"`
	Stores                                      string         `json:"stores"`
	StoresTags                                  []string       `json:"stores_tags"`
	Traces                                      string         `json:"traces"`
	TracesHierarchy                             []string       `json:"traces_hierarchy"`
	TracesTags                                  []string       `json:"traces_tags"`
	UniqueScansNumber                           int            `json:"unique_scans_n"`
	UpdateKey                                   string         `json:"update_key"`
}

type Ingredient struct {
	ID      string `json:"id"`
	Text    string `json:"text"`
	Rank    int    `json:"rank"`
	Percent string `json:percent`
}

type NutrientLevels struct {
	Salt         string `json:"salt"`
	Fat          string `json:"fat"`
	Sugars       string `json:"sugars"`
	SaturatedFat string `json:"saturated-fat"`
}

type Nutriments struct {
	Carbohydrates        json.Number `json:"carbohydrates"`
	Carbohydrates100G    json.Number `json:"carbohydrates_100g"`
	CarbohydratesServing json.Number `json:"carbohydrates_serving"`
	CarbohydratesUnit    string      `json:"carbohydrates_unit"`
	Energy               json.Number `json:"energy"`
	Energy100G           json.Number `json:"energy_100g"`
	EnergyServing        json.Number `json:"energy_serving"`
	EnergyUnit           string      `json:"energy_unit"`
	Fat                  json.Number `json:"fat"`
	Fat100G              json.Number `json:"fat_100g"`
	FatServing           json.Number `json:"fat_serving"`
	FatUnit              string      `json:"fat_unit"`
	Fiber                json.Number `json:"fiber"`
	Fiber100G            json.Number `json:"fiber_100g"`
	FiberServing         json.Number `json:"fiber_serving"`
	FiberUnit            string      `json:"fiber_unit"`
	Proteins             json.Number `json:"proteins"`
	Proteins100G         json.Number `json:"proteins_100g"`
	ProteinsServing      json.Number `json:"proteins_serving"`
	ProteinsUnit         string      `json:"proteins_unit"`
	Salt                 json.Number `json:"salt"`
	Salt100G             json.Number `json:"salt_100g"`
	SaltServing          json.Number `json:"salt_serving"`
	SaturatedFat         json.Number `json:"saturated-fat"`
	SaturatedFat100G     json.Number `json:"saturated-fat_100g"`
	SaturatedFatServing  json.Number `json:"saturated-fat_serving"`
	SaturatedFatUnit     string      `json:"saturated-fat_unit"`
	Sodium               json.Number `json:"sodium"`
	Sodium100G           json.Number `json:"sodium_100g"`
	SodiumServing        json.Number `json:"sodium_serving"`
	SodiumUnit           string      `json:"sodium_unit"`
	Sugars               json.Number `json:"sugars"`
	Sugars100G           json.Number `json:"sugars_100g"`
	SugarsServing        json.Number `json:"sugars_serving"`
	SugarsUnit           string      `json:"sugars_unit"`
	NutritionScoreFr     json.Number `json:"nutrition-score-fr"`
	NutritionScoreFr100G json.Number `json:"nutrition-score-fr_100g"`
	NutritionScoreUk     json.Number `json:"nutrition-score-uk"`
	NutritionScoreUk100G json.Number `json:"nutrition-score-uk_100g"`
}
