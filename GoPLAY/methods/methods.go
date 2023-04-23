package methods

type menuItem struct {
	name   string
	prices map[string]float64
}

type menu []menuItem

func (menuList menu) print() {

}

func(item *menuItem) add(){

	
}

var data = menu{
	{name: "latte", prices: map[string]float64{"small": 1.5}}
}

data.print()