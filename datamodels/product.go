package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID" rushbuy:"ID"`
	ProductName  string `json:"ProductName" sql:"productName" rushbuy:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" rushbuy:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" rushbuy:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" rushbuy:"ProductUrl"`
}
