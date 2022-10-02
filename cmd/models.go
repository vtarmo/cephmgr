package cmd

type Access struct {
	Ceph         string
	AccessKey    string
	AccessSecret string
}

type User struct {
	ID          string        `json:"user_id" url:"uid"`
	DisplayName string        `json:"display_name" url:"display-name"`
	Email       string        `json:"email" url:"email"`
	Keys        []UserKeySpec `json:"keys"`
	Caps        []UserCapSpec `json:"caps"`
	UserCaps    string        `url:"user-caps"`
}

type UserKeySpec struct {
	User      string `json:"user"`
	AccessKey string `json:"access_key" url:"access-key"`
	SecretKey string `json:"secret_key" url:"secret-key"`
}

type UserCapSpec struct {
	Type string `json:"type"`
	Perm string `json:"perm"`
}
