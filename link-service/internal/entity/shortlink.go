package entity

type ShortLink struct {
	ID          uint
	Owner       string
	OriginalUrl string
	ShortCode   string
	Clicks      int
}

func (sl *ShortLink) GetID() any {
	return sl.ID
}
