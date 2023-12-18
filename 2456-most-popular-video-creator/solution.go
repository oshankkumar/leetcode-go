package mostpopularvideocreator

import "container/heap"

type Video struct {
	Id      string
	Index   int
	Views   int
	Creator string
}

type VideoPQ []*Video

func (v VideoPQ) Len() int {
	return len(v)
}

func (v VideoPQ) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
	v[i].Index = i
	v[j].Index = j
}

func (v VideoPQ) Less(i, j int) bool {
	if v[i].Views == v[j].Views {
		return v[i].Id < v[j].Id
	}
	return v[i].Views > v[j].Views
}

func (v *VideoPQ) Push(x any) {
	video := x.(*Video)
	video.Index = len(*v)
	*v = append(*v, video)
}

func (v *VideoPQ) Pop() any {
	old := *v
	item := old[len(old)-1]
	old[len(old)-1] = nil
	*v = old[0 : len(old)-1]
	item.Index = -1
	return item
}

type PopularCreatorStore struct {
	creatorVideos map[string]*VideoPQ
	creatorViews  map[string]int
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	p := &PopularCreatorStore{
		creatorVideos: make(map[string]*VideoPQ),
		creatorViews:  make(map[string]int),
	}

	for i, creator := range creators {
		pq := p.creatorVideos[creator]
		if pq == nil {
			pq = &VideoPQ{}
			p.creatorVideos[creator] = pq
		}

		video := &Video{
			Id:      ids[i],
			Views:   views[i],
			Creator: creator,
		}
		heap.Push(pq, video)
		p.creatorViews[creator] += views[i]
	}

	var popularCreators []string
	for creator, views := range p.creatorViews {
		if len(popularCreators) == 0 {
			popularCreators = append(popularCreators, creator)
			continue
		}

		if views > p.creatorViews[popularCreators[0]] {
			popularCreators = popularCreators[:0]
			popularCreators = append(popularCreators, creator)
			continue
		}

		if views == p.creatorViews[popularCreators[0]] {
			popularCreators = append(popularCreators, creator)
		}
	}

	var popularVideos [][]string

	for _, c := range popularCreators {
		video := (*(p.creatorVideos[c]))[0]
		popularVideos = append(popularVideos, []string{c, video.Id})
	}

	return popularVideos
}
