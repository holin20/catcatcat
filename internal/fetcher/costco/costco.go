package costco

import (
	"os"

	"github.com/holin20/catcatcat/pkg/ezgo"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

type ItemModel struct {
	ItemId     string
	ProductId  string
	CategoryId string
	Available  bool
	Price      float64
}

func FetchItem(inventroyUrl, priceUrl string) (float64, bool, error) {
	priceArgs, invArgs := ezgo.Await2(
		ezgo.Async2(ezgo.Bind2_2(fetchJsonPath, priceUrl, "finalOnlinePrice")),
		ezgo.Async2(ezgo.Bind2_2(fetchJsonPath, inventroyUrl, "invAvailable")),
	)

	priceResult, priceErr := priceArgs.Unpack()
	hasInvResult, invErr := invArgs.Unpack()

	var price float64
	var hasInv bool
	if priceResult != nil {
		price = priceResult.Float()
	}
	if hasInvResult != nil {
		hasInv = hasInvResult.Bool()
	}

	return price, hasInv, ezgo.If(
		ezgo.IsErr(priceErr) && ezgo.IsErr(invErr),
		ezgo.NewCause(priceErr, "fetchJsonPath.finalOnlinePrice"),
		nil,
	)
}

func FetchItemModel(
	scope *ezgo.Scope,
	name string,
	itemId string,
	categoryId string,
	productId string,
	queryStringPatch string,
) (*ItemModel, error) {
	scope = scope.WithLogger(scope.GetLogger().Named("CostcoFetcher"))

	// fetch price
	priceUrl := buildGetContractPriceUrl(itemId, categoryId, productId, queryStringPatch)
	scope.GetLogger().Info("GetContractPriceUrl", zap.String("priceUrl", priceUrl), zap.String("name", name))
	priceResult, err := fetchJsonPath(priceUrl, "finalOnlinePrice")
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "fetchJsonPath.finalOnlinePrice."+priceUrl)
	}

	// fetch inventory
	inventroyUrl := builGetInventoryDetailUrl(itemId, categoryId, productId, queryStringPatch)
	scope.GetLogger().Info("GetInventoryDetailUrl", zap.String("inventroyUrl", inventroyUrl), zap.String("name", name))
	hasInvResult, err := fetchJsonPath(inventroyUrl, "invAvailable")
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "fetchJsonPath.invAvailable."+inventroyUrl)
	}

	return &ItemModel{
		ItemId:     itemId,
		CategoryId: categoryId,
		ProductId:  productId,
		Price:      priceResult.Float(),
		Available:  hasInvResult.Bool(),
	}, nil
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
		return nil, ezgo.NewCausef(err, "ExtractJsonPath(%s, %s)", ezgo.FirstNChars(body, 200), path)
	}

	return result, nil
}

func getCookieString() string {
	return os.Getenv("CATCATCAT_COSTCO_COOKIE")
}

func buildGetContractPriceUrl(
	itemId string,
	catalogId string,
	productId string,
	queryStringPatch string,
) string {
	// e.g. https://www.costco.com/AjaxGetContractPrice?itemId=3074457345620439817&catalogId=10701&productId=3074457345620439815
	return ezgo.NewHttpsUrl("www.costco.com").
		WithPath("AjaxGetContractPrice").
		WithQueryParam("itemId", itemId).
		WithQueryParam("catalogId", catalogId).
		WithQueryParam("productId", productId).
		WithQueryStringPatch(queryStringPatch).
		String()
}

func builGetInventoryDetailUrl(
	itemId string,
	catalogId string,
	productId string,
	queryStringPatch string,
) string {
	// e.g. https://www.costco.com/AjaxGetContractPrice?itemId=3074457345620439817&catalogId=10701&productId=3074457345620439815
	return ezgo.NewHttpsUrl("www.costco.com").
		WithPath("AjaxGetInventoryDetail").
		WithQueryParam("itemId", itemId).
		WithQueryParam("catalogId", catalogId).
		WithQueryParam("productId", productId).
		WithQueryStringPatch(queryStringPatch).
		String()
}
