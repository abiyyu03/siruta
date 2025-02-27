package constant

type Constant struct{}

func (c *Constant) RomanMonth() map[string]string {
	return map[string]string{
		"1":  "I",
		"2":  "II",
		"3":  "III",
		"4":  "IV",
		"5":  "V",
		"6":  "VI",
		"7":  "VII",
		"8":  "VIII",
		"9":  "IX",
		"10": "X",
		"11": "XI",
		"12": "XII",
	}
}

func (c *Constant) ResponseMessage() map[string]string {
	return map[string]string{
		"success":              "Berhasil",
		"error":                "Gagal",
		"not_found":            "Data tidak ditemukan",
		"unauthorized":         "Akses ditolak",
		"invalid_data":         "Data yang dimasukkan salah",
		"already_exists":       "Data sudah ada",
		"duplicate_data":       "Data sudah ada",
		"internal_error":       "Terjadi kesalahan internal",
		"unprocessable_entity": "Entitas yang dimasukkan tidak dapat diproses",
		"validation_failed":    "Validasi gagal",
		"already_registered":   "Anda sudah terdaftar",
		"not_registered":       "Anda belum terdaftar",
		"email_already_exists": "Email sudah terdaftar",
	}
}
