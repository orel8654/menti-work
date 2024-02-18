package types

type User struct {
	IDKey    int    `db:"id_key"`
	UUID     string `db:"uuid"`
	Password string `db:"password"`
	Username string `db:"username"`
}
