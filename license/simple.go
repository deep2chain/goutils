package license

import (
	gu "github.com/deep2chain/goutils"
)

func IsLicenseValid(url string) bool {
	content, _ := gu.GetWebSite(url)
	if content != nil {
		return true
	}
	return false
}
