package piscine

type food struct {
	burger  int
	chips   int
	nuggets int
}

func FoodDeliveryTime(order string) int {
	belale := food{
		burger:  15,
		chips:   10,
		nuggets: 12,
	}
	if order == "burger" {
		return belale.burger
	} else if order == "chips" {
		return belale.chips
	} else if order == "nuggets" {
		return belale.nuggets
	}
	return 404
}
