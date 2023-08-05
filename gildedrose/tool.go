package gildedrose

func isSulfuras(item *Item) bool {
	return item.Name == "Sulfuras, Hand of Ragnaros"
}
func isBackstage(item *Item) bool {
	return item.Name == "Backstage passes to a TAFKAL80ETC concert"
}

func isAgeBrie(item *Item) bool {
	return item.Name == "Aged Brie"
}

func isConj(item *Item) bool {
	return item.Name == "Conjured"
}

func checkItem(item *Item) int {
	if isSulfuras(item) {
		return 0
	}
	if isAgeBrie(item) {
		return 1
	}
	if isBackstage(item) {
		return 2
	}
	if isConj(item) {
		return 3
	}
	return 4
}

func updateSulfuras(item *Item) {
	return
}
func updateAgeBrie(item *Item) {
	if item.Quality < 50 {
		item.addQuality()
	}
	if item.SellIn <= 0 && item.Quality < 50 {
		item.addQuality()
	}
	item.reduceSellIn()
}
func updateBackstage(item *Item) {
	if item.Quality < 50 {
		item.addQuality()
		if item.SellIn < 11 && item.Quality < 50 {
			item.addQuality()
		}
		if item.SellIn < 6 && item.Quality < 50 {
			item.addQuality()
		}
	}
	item.reduceSellIn()
	if item.SellIn < 0 {
		item.Quality = 0
	}
}
func updateNormal(item *Item) {
	if item.Quality > 0 {
		item.reduceQuality()
	}
	item.SellIn = item.SellIn - 1
	if item.SellIn < 0 {
		if item.Quality > 0 {
			item.reduceQuality()
		}
	}
}

func updateConj(item *Item) {
	if item.SellIn <= 0 {
		item.Quality = max(0, item.Quality-4)
	} else {
		item.Quality = max(0, item.Quality-2)
	}
	item.SellIn = item.SellIn - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
