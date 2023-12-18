package foodratingsystem

import "container/heap"

type FoodItem struct {
	rating  int
	name    string
	index   int
	cuisine string
}

type FoodItemPQ []*FoodItem

func (f FoodItemPQ) Len() int {
	return len(f)
}

func (f FoodItemPQ) Less(i, j int) bool {
	if f[i].rating == f[j].rating {
		return f[i].name < f[j].name
	}

	return f[i].rating > f[j].rating
}

func (f FoodItemPQ) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
	f[i].index = i
	f[j].index = j
}

func (f *FoodItemPQ) Push(x any) {
	item := x.(*FoodItem)
	item.index = len(*f)
	*f = append(*f, item)
}

func (f *FoodItemPQ) Pop() any {
	old := *f
	item := old[len(old)-1]
	old[len(old)-1] = nil
	*f = old[:len(old)-1]
	item.index = -1
	return item
}

type FoodRatings struct {
	cuisines map[string]*FoodItemPQ
	foods    map[string]*FoodItem
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	f := FoodRatings{
		cuisines: make(map[string]*FoodItemPQ),
		foods:    make(map[string]*FoodItem),
	}

	for i, foodName := range foods {
		pq := f.cuisines[cuisines[i]]
		if pq == nil {
			pq = &FoodItemPQ{}
			f.cuisines[cuisines[i]] = pq
		}

		item := &FoodItem{name: foodName, rating: ratings[i], cuisine: cuisines[i]}
		f.foods[foodName] = item
		heap.Push(pq, item)
	}

	return f
}

func (f *FoodRatings) ChangeRating(food string, newRating int) {
	foodItem := f.foods[food]
	foodItem.rating = newRating
	heap.Fix(f.cuisines[foodItem.cuisine], foodItem.index)
}

func (f *FoodRatings) HighestRated(cuisine string) string {
	foodPQ := f.cuisines[cuisine]
	return (*foodPQ)[0].name
}
