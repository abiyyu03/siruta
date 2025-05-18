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

type ErrorMessage struct {
	Message string `json:"message"`
	Clue    string `json:"clue"`
}

var Errors = map[string]ErrorMessage{
	"Success":               {Message: "Berhasil", Clue: "Operasi berhasil dilakukan."},
	"Error":                 {Message: "Gagal", Clue: "Terjadi kesalahan, coba lagi nanti."},
	"NotFound":              {Message: "Data tidak ditemukan", Clue: "Pastikan data yang dicari sudah benar atau tersedia."},
	"UserNotFound":          {Message: "Data Pengguna tidak ditemukan", Clue: "Pastikan pengguna yang dicari sudah terdaftar di sistem"},
	"LetterRejected":        {Message: "Surat ditolak", Clue: "Pastikan pengajuan anda memiliki data yang akurat atau komunikasikan dengan pengurus RT."},
	"Unauthorized":          {Message: "Akses ditolak", Clue: "Pastikan Anda memiliki izin yang sesuai."},
	"AccountInputError":     {Message: "Email atau Password anda salah", Clue: "Pastikan Anda memasukannya dengan benar"},
	"InvalidData":           {Message: "Data yang dimasukkan salah", Clue: "Periksa kembali format dan isi data yang dikirim."},
	"InvalidReferalCode":    {Message: "Referal Code salah", Clue: "Periksa kembali referal code yang dikirim"},
	"AlreadyExists":         {Message: "Data sudah ada", Clue: "Coba gunakan data yang berbeda atau edit data yang ada."},
	"DuplicateData":         {Message: "Data sudah ada", Clue: "Hindari memasukkan data yang sama lebih dari satu kali."},
	"InternalError":         {Message: "Terjadi kesalahan internal", Clue: "Silakan coba lagi nanti atau hubungi admin."},
	"UnprocessableEntity":   {Message: "Entitas yang dimasukkan tidak dapat diproses", Clue: "Pastikan semua data sudah sesuai format yang diharapkan."},
	"ValidationFailed":      {Message: "Validasi gagal", Clue: "Periksa kembali data yang dikirim dan pastikan sudah sesuai aturan."},
	"AlreadyRegistered":     {Message: "Anda sudah terdaftar", Clue: "Gunakan akun yang sudah ada untuk masuk."},
	"NotRegistered":         {Message: "Anda belum terdaftar", Clue: "Silakan daftar terlebih dahulu sebelum mengakses layanan ini."},
	"EmailAlreadyExists":    {Message: "Email sudah terdaftar", Clue: "Gunakan email lain atau login dengan email ini."},
	"InvalidToken":          {Message: "Token tidak valid", Clue: "Pastikan Anda menggunakan token yang benar atau lakukan login ulang."},
	"TokenExpired":          {Message: "Token sudah kedaluwarsa", Clue: "Silakan login ulang untuk mendapatkan token baru."},
	"InvalidPassword":       {Message: "Password yang dimasukkan salah", Clue: "Coba periksa kembali password Anda atau lakukan reset password."},
	"NotMatchPassword":      {Message: "Password konfirmasi tidak cocok", Clue: "Coba periksa kembali inputan anda"},
	"InvalidRole":           {Message: "Role yang dimasukkan salah", Clue: "Pastikan Anda menggunakan role yang sesuai."},
	"RoleNotFound":          {Message: "Role yang dimasukkan tidak ditemukan", Clue: "Cek kembali role yang digunakan atau hubungi admin."},
	"InvalidRolePermission": {Message: "Anda tidak memiliki izin untuk mengakses halaman ini", Clue: "Pastikan akun Anda memiliki hak akses yang sesuai."},
	"InvalidRequestMethod":  {Message: "Metode permintaan yang dimasukkan salah", Clue: "Gunakan metode yang sesuai (GET, POST, PUT, DELETE, dll.)."},
	"InvalidDateFormat":     {Message: "Format tanggal yang dimasukkan salah", Clue: "Gunakan format yang benar, misalnya YYYY-MM-DD."},
	"InvalidNumber":         {Message: "Angka yang dimasukkan salah", Clue: "Pastikan angka yang dimasukkan sesuai dengan format yang benar."},
	"EmailQueryRequired":    {Message: "Parameter query email diperlukan", Clue: "Pastikan untuk menyertakan parameter 'email' dalam permintaan Anda."},
}
