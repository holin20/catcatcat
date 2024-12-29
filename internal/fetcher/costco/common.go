package costco

import (
	"github.com/holin20/catcatcat/pkg/ezgo"
)

func FetchMacbookPrice() (float64, error) {
	url := "https://www.costco.com/AjaxGetContractPrice?itemId=3074457345620577642&catalogId=10701&productId=3074457345620577640&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false"
	body, err := ezgo.NewHttpClient().WithDefaultUserAgent().Get(url)
	if ezgo.IsErr(err) {
		return 0, ezgo.NewCause(err, "HttpCall")
	}

	price, err := ezgo.GetFloatFromJSONPath(string(body), "finalOnlinePrice")
	if ezgo.IsErr(err) {
		return 0, ezgo.NewCause(err, "GetFloatFromJSONPath")
	}

	return price, nil
}
