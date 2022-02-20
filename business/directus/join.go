package directus

import "fmt"

type (
	Email struct {
		Email string `json:"email" validate:"required,email"`
	}
	Password struct {
		Password string `json:"password" validate:"required,max=32,min=6"`
	}
	BaseUser struct {
		Email
		Password
	}
	User struct {
		BaseUser
		Role string `json:"role"`
	}
	Member struct {
		StudentID   string `json:"学号" validate:"required,max=12"`
		RealName    string `json:"姓名" validate:"required,max=8"`
		Major       int    `json:"专业" validate:"required,lt=2000"` // 见CMS中的专业列表
		Sexual      string `json:"性别" validate:"required,oneof=male female other"`
		Grade       int    `json:"年级" validate:"required,gt=2001"`
		Phone       string `json:"手机" validate:"required,max=11,min=6"`
		PhoneShort  string `json:"短号,omitempty"`
		Wechat      string `json:"微信" validate:"required,min=6,max=20"`
		Group       string `json:"职能" validate:"required"` // 数组形式的职能组 e.g. ["技术开发"]
		JoinDate    string `json:"加入日期,omitempty"`
		Dormitory   string `json:"宿舍,omitempty"`
		NewBirthday string `json:"公历生日,omitempty"`
		OldBirthday string `json:"农历生日,omitempty"`
	}
	UserMember struct {
		User
		Member
	}
)

func (u *User) Create() bool {
	resp, err := client.R().
		SetBody(*u).
		Post("/users")
	fmt.Println(resp.RawResponse, resp.String())
	return err == nil && resp.IsSuccess()
}

func (m *Member) Create() bool {
	resp, err := client.R().
		SetBody(*m).
		Post("/items/directus_users")
	fmt.Println(resp.RawResponse, resp.String())
	return err == nil && resp.IsSuccess()
}
