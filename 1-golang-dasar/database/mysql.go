package database

var connection string

func init() { // function yg langsung di-eksekusi saat program dijalankan
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
