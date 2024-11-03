package methods

type MenuItem struct {
	Name   string
	Prices map[string]float64
}

type Menu []MenuItem

func (item *MenuItem) add() {

}

var mydata Menu = Menu{
	MenuItem{
		Name: "yam",
		Prices: map[string]float64{
			"default": 1.24,
		},
	},
}
