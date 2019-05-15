package helper

type Pagination struct {
	Page  uint `form:"p"`
	Limit uint `form:"l"`
}

func (pagination *Pagination) GetPage() uint {
	if pagination.Page == 0 {
		return 1
	}
	if pagination.Page > 100 {
		return 100
	}
	return pagination.Page
}

func (pagination *Pagination) GetLimit() uint {
	if pagination.Limit == 0 || pagination.Limit > 5 {
		return 5
	}
	return pagination.Limit
}

func (pagination *Pagination) GetOffSet() uint {
	page := pagination.GetPage()
	limit := pagination.GetLimit()
	offset := (page - 1) * limit
	return offset
}
