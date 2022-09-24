package requests

type HeaderItem struct {
	Header
	ToItem `json:"to_CreditMemoRequest"`
}

type ToItem struct {
	Results []Item `json:"results"`
}
