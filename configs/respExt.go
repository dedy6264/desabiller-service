package configs

var (
	RC_SUCCESS               = []string{"00", "SUCCESS"}
	RC_PENDING               = []string{"02", "PENDING"}
	RC_FAILED                = []string{"03", "FAILED"}
	RC_INQUIRY_SUCCESS       = []string{"04", "INQUIRY SUCCESS"}
	RC_INQUIRY_FAILED        = []string{"11", "INQUIRY FAILED"}
	RC_EXPIRED               = []string{"05", "EXPIRED"}
	RC_SUSPECT               = []string{"06", "SUSPECT"}
	RC_INVALID_PARAM         = []string{"39", "INVALID REQUEST PARAMETER"}
	RC_SYSTEM_ERROR          = []string{"07", "SYSTEM ERROR"}
	RC_TERMINATED            = []string{"08", "TRANSACTION IS TERMINATE"}
	RC_DUPLICATE_REF         = []string{"09", "DUPLICATE PARTNER REFERENCE NO"}
	RC_PRODUCT_NOT_FOUND     = []string{"10", "PRODUCT CODE NOT EXISTS"}
	RC_JWT_EXPIRED           = []string{"401", "JWT EXPIRED"}
	RC_INVALID_PARTNER       = []string{"12", "INVALID PARTNER ID"}
	RC_INVALID_SIGNATURE     = []string{"13", "INVALID DIGITAL SIGNATURE"}
	RC_INVALID_CUSTOMER_ID   = []string{"14", "INVALID CUSTOMER ID"}
	RC_INVALID_TRANSACTION   = []string{"15", "INVALID TRANSACTION"}
	RC_CUSTOMER_INACTIVE     = []string{"16", "INACTIVE CUSTOMER/ACCOUNT"}
	RC_ALREADY_PAID          = []string{"17", "PAYMENT ALREADY PAID"}
	RC_INQUIRY_ALREADY_PAID  = []string{"18", "INQUIRY ALREADY PAID"}
	RC_PERIOD_LIMIT          = []string{"19", "PERIOD OUT OF MAX LIMIT"}
	RC_INQUIRY_ONLY          = []string{"20", "INQUIRY ONLY/PURC_HASE DATA NOT AVAILABLE"}
	RC_TRANSACTION_ABORTED   = []string{"21", "TRANSACTION IS ABORTED"}
	RC_TRANSACTION_NOT_FOUND = []string{"22", "TRANSACTION IS NOT FOUND"}
	RC_BILLER_REJECTED       = []string{"24", "TRANSACTION CAN'T BE PROCESSED BY BILLER/OPERATOR"}
	RC_ALREADY_SUCCESS       = []string{"25", "TRANSACTION ALREADY SUCCESS ON 1X24 HOUR"}
	RC_CROSS_REGION          = []string{"26", "TRANSACTION CROSS CLUSTER/REGION"}
	RC_DUPLICATE_TRANSACTION = []string{"27", "DUPLICATE TRANSACTION"}
	RC_BILLING_NOT_FOUND     = []string{"28", "BILLING INFORMATION CAN'T BE FOUND/ALREADY PAID OR TOTAL KWH OUT OF MAX LIMIT"}
	RC_PRODUCT_DISRUPTION    = []string{"29", "PRODUCT DISRUPTION"}
	RC_PRODUCT_CLOSED        = []string{"30", "PRODUCT CLOSED"}
	RC_EXCESSIVE_BALANCE     = []string{"36", "EXCESSIVE BALANCE"}
	RC_ACCOUNT_BLOCKED       = []string{"37", "CUSTOMER ACCOUNT IS BLOCKED"}
	RC_INVALID_AMOUNT        = []string{"38", "INVALID AMOUNT (AMOUNT BELOW MINIMUM)"}
	RC_TIMEOUT_RETRY         = []string{"40", "TRANSACTION TIME LIMIT REACHED, PLEASE WAIT A MOMENT."}
	RC_BILLER_DISRUPTION     = []string{"96", "BILLER PRINCIPAL DISRUPTION"}
	RC_SYSTEM_CUTOFF         = []string{"97", "SYSTEM CUT-OFF"}

	RC_VALIDATION_FAILED        = []string{"01", "Validasi gagal"}
	RC_FAILED_USER_EXISTING     = []string{"31", "coba dengan username/email lain"}
	RC_FAILED_USER_NOT_FOUND    = []string{"32", "User tidak ditemukan"}
	RC_FAILED_WRONG_PWD_USRNAME = []string{"33", "password/username salah"}
	RC_FAILED_WRONG_OTP         = []string{"34", "OTP tidak valid	OTP salah / expired"}

	RC_FAILED_DB_NOT_FOUND = []string{"82", "Not found"}
)

// Code	Message	Keterangan
// 00	Sukses	Request berhasil
// 		Format input salah, field kosong, email invalid
// 02	Username/email sudah terdaftar	Saat registrasi
// 03	User tidak ditemukan	Untuk login atau reset password
// 04	Password salah	Login gagal
// 05	OTP tidak valid	OTP salah / expired
// 06	Akses ditolak	Token tidak valid / expired
// 07	Akun tidak aktif	Akun terblokir / belum aktivasi email
// 08	Session habis	User perlu login ulang
// 09	Tidak diizinkan	User tidak punya izin untuk aksi tertentu
// 10	Terlalu sering	Rate limit, pencegahan brute-force
// 99	Sistem error	Server error, unexpected exception
