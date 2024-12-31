package costco

import (
	"github.com/holin20/catcatcat/pkg/ezgo"
)

func FetchItemPrice(url string) (float64, error) {
	body, err := ezgo.NewHttpClient().WithDefaultUserAgent().Get(url)
	if ezgo.IsErr(err) {
		return 0, ezgo.NewCause(err, "HttpCall")
	}

	priceField := "finalOnlinePrice"
	price, err := ezgo.GetFloatFromJSONPath(string(body), priceField)
	if ezgo.IsErr(err) {
		return 0, ezgo.NewCausef(err, "GetFloatFromJSONPath(%s, %s)", body, priceField)
	}

	return price, nil
}
