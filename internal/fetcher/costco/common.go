package costco

import (
	"github.com/holin20/catcatcat/pkg/ezhttp"
	"github.com/holin20/catcatcat/pkg/ezjson"
	"github.com/holin20/catcatcat/pkg/gen"
)

func FetchMacbookPrice() (float64, error) {
	url := "https://www.costco.com/AjaxGetContractPrice?itemId=3074457345620577642&catalogId=10701&productId=3074457345620577640&isFrozenItem=false&isRegionalPdt=false&isBundleItem=false"
	body, err := ezhttp.NewHttpClient().WithDefaultUserAgent().Get(url)
	if gen.IsErr(err) {
		return 0, err
	}

	price, err := ezjson.GetFloatFromJSONPath(string(body), "finalOnlinePrice")
	if gen.IsErr(err) {
		return 0, err
	}

	return price, nil
}
