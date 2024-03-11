package types

type UserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserPayloadCreat struct {
	Username string `db:"username"`
	UUID     string `db:"uuid"`
	Password string `db:"password"`
}

type User struct {
	IDKey    int    `db:"id_key"`
	UUID     string `db:"uuid"`
	Password string `db:"password"`
	Username string `db:"username"`
}

type UserToken struct {
	Token string `json:"token" db:"token"`
	IDKey int `json:"id_key" db:"id_key"`
}

type ResponseError map[string]string

type ResponseLogin map[string]string