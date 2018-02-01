package domain

import "errors"

const (
	MaterialTypeSeedCode                = "MATERIAL_SEED"
	MaterialTypeGrowingMediumCode       = "MATERIAL_GROWING_MEDIUM"
	MaterialTypeAgrochemicalCode        = "MATERIAL_AGROCHEMICAL"
	MaterialTypeLabelAndCropSupportCode = "MATERIAL_LABEL_AND_CROP_SUPPORT"
	MaterialTypeSeedingContainerCode    = "MATERIAL_SEEDING_CONTAINER"
	MaterialTypePostHarvestSupplyCode   = "MATERIAL_POST_HARVEST_SUPPLY"
	MaterialTypeOtherCode               = "MATERIAL_OTHER"
)

type MaterialType interface {
	Code() string
}

type MaterialTypeSeed struct {
	PlantType PlantType
}

func (mt MaterialTypeSeed) Code() string {
	return MaterialTypeSeedCode
}

func CreateMaterialTypeSeed(plantType string) (MaterialTypeSeed, error) {
	pt := GetPlantType(plantType)
	if pt == (PlantType{}) {
		return MaterialTypeSeed{}, errors.New("options wrong")
	}

	return MaterialTypeSeed{pt}, nil
}

type PlantType struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

const (
	PlantTypeVegetable = "VEGETABLE"
	PlantTypeFruit     = "FRUIT"
	PlantTypeHerb      = "HERB"
	PlantTypeFlower    = "FLOWER"
	PlantTypeTree      = "TREE"
)

func PlantTypes() []PlantType {
	return []PlantType{
		{Code: PlantTypeVegetable, Label: "Vegetable"},
		{Code: PlantTypeFruit, Label: "Fruit"},
		{Code: PlantTypeHerb, Label: "Herb"},
		{Code: PlantTypeFlower, Label: "Flower"},
		{Code: PlantTypeTree, Label: "Tree"},
	}
}

func GetPlantType(code string) PlantType {
	for _, v := range PlantTypes() {
		if v.Code == code {
			return v
		}
	}

	return PlantType{}
}

type MaterialTypeAgrochemical struct {
	ChemicalType ChemicalType
}

func (mt MaterialTypeAgrochemical) Code() string {
	return MaterialTypeAgrochemicalCode
}

const (
	ChemicalTypeDisinfectant = "DISINFECTANT"
	ChemicalTypeFertilizer   = "FERTILIZER"
	ChemicalTypeHormone      = "HORMONE"
	ChemicalTypeManure       = "MANURE"
	ChemicalTypePesticide    = "PESTICIDE"
)

type ChemicalType struct {
	Code  string
	Label string
}

func ChemicalTypes() []ChemicalType {
	return []ChemicalType{
		{Code: ChemicalTypeDisinfectant, Label: "Disinfectant and Sanitizer"},
		{Code: ChemicalTypeFertilizer, Label: "Fertilizer"},
		{Code: ChemicalTypeHormone, Label: "Hormone and Growth Agent"},
		{Code: ChemicalTypeManure, Label: "Manure"},
		{Code: ChemicalTypePesticide, Label: "Pesticide"},
	}
}

func GetChemicalType(code string) ChemicalType {
	for _, v := range ChemicalTypes() {
		if v.Code == code {
			return v
		}
	}

	return ChemicalType{}
}

func CreateMaterialTypeAgrochemical(chemicalType string) (MaterialTypeAgrochemical, error) {
	ct := GetChemicalType(chemicalType)
	if ct == (ChemicalType{}) {
		return MaterialTypeAgrochemical{}, errors.New("options wrong")
	}

	return MaterialTypeAgrochemical{ct}, nil
}