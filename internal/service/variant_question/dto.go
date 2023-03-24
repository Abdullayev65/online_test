package variant_question_srvc

type Create struct {
	TopicID   int `json:"topic_id"`
	Count     int `json:"count"`
	VariantID int `json:"-"`
}

type Filter struct {
	Limit          *int
	Offset         *int
	Order          *string
	AllWithDeleted bool
	OnlyDeleted    bool
	VariantID      *int `json:"-"`
}
