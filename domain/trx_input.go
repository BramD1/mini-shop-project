package domain

type DetailTrxInput struct {
    ProductID uint `json:"product_id"`
    Kuantitas uint `json:"kuantitas"`
}

type TrxInput struct {
    MethodBayar  string           `json:"method_bayar"`
    AlamatKirim  uint             `json:"alamat_kirim"`
    DetailTrx    []DetailTrxInput `json:"detail_trx"`
}