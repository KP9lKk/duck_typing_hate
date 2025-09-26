package entity

type ShortLink struct {
	ID          uint
	Owner       string
	OriginalUrl string
	ShortCode   string
	Clicks      int
}
