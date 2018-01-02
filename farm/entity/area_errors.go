package entity

const (
	AreaErrorNameEmptyCode = iota
	AreaErrorNameNotEnoughCharacterCode
	AreaErrorNameExceedMaximunCharacterCode
	AreaErrorNameAlphanumericOnlyCode

	AreaErrorSizeEmptyCode

	AreaErrorTypeEmptyCode
	AreaErrorInvalidAreaTypeCode

	AreaErrorLocationEmptyCode
	AreaErrorInvalidAreaLocationCode
)

// AreaError is a custom error from Go built-in error
type AreaError struct {
	Code int
}

func (e AreaError) Error() string {
	switch e.Code {
	case AreaErrorNameEmptyCode:
		return "Area name is required."
	case AreaErrorNameNotEnoughCharacterCode:
		return "Not enough character on Area Name"
	case AreaErrorNameExceedMaximunCharacterCode:
		return "Area name cannot more than 100 characters"
	case AreaErrorNameAlphanumericOnlyCode:
		return "Area name should be alphanumeric"
	case AreaErrorSizeEmptyCode:
		return "Area size cannot be empty"
	case AreaErrorTypeEmptyCode:
		return "Area type cannot be empty"
	case AreaErrorInvalidAreaTypeCode:
		return "Area type is invalid"
	case AreaErrorLocationEmptyCode:
		return "Area location cannot be empty"
	case AreaErrorInvalidAreaLocationCode:
		return "Area location is invalid"
	default:
		return "Unrecognized Reservoir Error Code"
	}
}