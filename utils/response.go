package utils

type Response struct {
    Status  bool        `json:"status"`
    Message string      `json:"message"`
    Errors  interface{} `json:"errors"`
    Data    interface{} `json:"data"`
}

// Add to utils/response.go
type UserResponse struct {
    ID           uint   `json:"id"`
    Nama         string `json:"nama"`
    NoTelp       string `json:"no_telp"`
    TanggalLahir string `json:"tanggal_lahir"`
    Pekerjaan    string `json:"pekerjaan"`
    Email        string `json:"email"`
    ProvinsiID   string `json:"id_provinsi"`
    KotaID       string `json:"id_kota"`
    IsAdmin      bool   `json:"is_admin"`
    Token        string `json:"token,omitempty"`
}

func SuccessResponse(message string, data interface{}) Response {
    return Response{
        Status:  true,
        Message: message,
        Errors:  nil,
        Data:    data,
    }
}

func ErrorResponse(message string, errors interface{}) Response {
    return Response{
        Status:  false,
        Message: message,
        Errors:  errors,
        Data:    nil,
    }
}