package vectors

type Player struct {
	Level      int
	Kost       float64
	Rank       int
	RankPoints int
}

// Convert an interface to a float vector
func ConvertPlayerToVector(player Player) []float32 {
	return []float32{float32(player.Level), float32(player.Kost), float32(player.Rank), float32(player.RankPoints)}
}

type numeric interface {
	int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func numericSliceToFloat32Vector[T numeric](slice []T) []float32 {
	float32Vector := make([]float32, len(slice))

	for i, val := range slice {
		float32Vector[i] = float32(val)
	}

	return float32Vector
}
