package directus

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/mix-go/dotenv"
)

var directusUrl string
var directusToken string

var client *resty.Client

func init() {
	_ = dotenv.Load(".env")
	directusUrl = dotenv.Getenv("DIRECTUS_URL").String()
	directusToken = dotenv.Getenv("DIRECTUS_TOKEN").String()

	client = resty.New()
	client.SetBaseURL(directusUrl)
	client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", directusToken))
	client.SetHeader("User-Agent", "ZCST XMT OA/1.0")
}
