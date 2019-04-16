package orderset

type OrderSet struct {
	set map[interface{}]int
	order []interface{}
}


func NewIntegerOrderSet() OrderSet {
	return OrderSet{set: map[interface{}]int{},order: []interface{}{}}
}


func (s *OrderSet) Add(element interface{}){
	if _, exist := s.set[element]; !exist {
		// 透過目前陣列內的元素個數判斷將要放入的索引位置
		index := len(s.order)
		s.set[element] = index
		s.order = append(s.order,element)
	}
}

func (s *OrderSet) Del(element interface{}){
	if index, exist := s.set[element]; exist {
		// 取得索引刪除元素
		delete(s.set,element)
		s.order = append(s.order[:index], s.order[index+1:]...)
	}
}

func (s *OrderSet) Contains(element interface{})(bool){
	_, exist := s.set[element]
	return exist
}

func (s *OrderSet) ToSlice()(slice []interface{}){
	return s.order
}

func (s *OrderSet) Count() int {
	return len(s.order)
}