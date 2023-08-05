package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func (item *Item) reduceQuality() {
	item.Quality = item.Quality - 1
}

func (item *Item) addQuality() {
	item.Quality = item.Quality + 1
}

func (item *Item) reduceSellIn() {
	item.SellIn = item.SellIn - 1
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		UpdateItem(item)
	}
}

func UpdateItem(item *Item) {
	switch checkItem(item) {
	case 0:
		updateSulfuras(item)
	case 1:
		updateAgeBrie(item)
	case 2:
		updateBackstage(item)
	case 3:
		updateConj(item)
	default:
		updateNormal(item)
	}
}
