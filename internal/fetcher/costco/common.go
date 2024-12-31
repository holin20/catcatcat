package costco

import (
	"os"

	"github.com/holin20/catcatcat/pkg/ezgo"
	"github.com/tidwall/gjson"
)

func FetchItem(inventroyUrl, priceUrl string) (float64, bool, error) {
	priceResult, errP := fetchJsonPath(priceUrl, "finalOnlinePrice")
	hasInvResult, errI := fetchJsonPath(inventroyUrl, "invAvailable")

	var price float64
	var hasInv bool
	if priceResult != nil {
		price = priceResult.Float()
	}
	if hasInvResult != nil {
		hasInv = hasInvResult.Bool()
	}

	return price, hasInv, ezgo.If(
		ezgo.IsErr(errI) && ezgo.IsErr(errP),
		ezgo.NewCause(errI, "fetchJsonPath.finalOnlinePrice"),
		nil,
	)
}

func fetchJsonPath(url string, path string) (*gjson.Result, error) {
	body, err := ezgo.NewHttpClient().
		WithDefaultUserAgent().
		SetCookieString(getCookieString()).
		Get(url)
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCausef(err, "HttpCall(%s)", url)
	}

	result, err := ezgo.ExtractJsonPath(string(body), path)
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCausef(err, "ExtractJsonPath(%s, %s)", body, path)
	}

	return result, nil
}

func getCookieString() string {
	return os.Getenv("CATCATCAT_COSTCO_COOKIE")
}
