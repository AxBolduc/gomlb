package mlb

type PlayerHittingStats struct {
	Type struct {
		DisplayName string
	}
	Group struct {
		DisplayName string
	}
	Splits []HittingStatsSplit
}
type PlayerPitchingStats struct {
	Type struct {
		DisplayName string
	}
	Group struct {
		DisplayName string
	}
	Splits []PitchingStatsSplit
}
type PlayerFieldingStats struct {
	Type struct {
		DisplayName string
	}
	Group struct {
		DisplayName string
	}
	Splits []FieldingStatsSplit
}
