package costco

import (
	"fmt"
	"os"
	"time"

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

func FetchItemModel(
	scope *ezgo.Scope,
	httpClient *ezgo.HttpClient,
	name string,
	itemId string,
	categoryId string,
	productId string,
	queryStringPatch string,
) (*ItemModel, error) {
	scope = scope.WithLogger(scope.GetLogger().Named("CostcoFetcher"))

	priceUrl := buildGetContractPriceUrl(itemId, categoryId, productId, queryStringPatch)
	scope.GetLogger().Info("GetContractPriceUrl", zap.String("priceUrl", priceUrl), zap.String("name", name))

	inventroyUrl := builGetInventoryDetailUrl(itemId, categoryId, productId, queryStringPatch)
	scope.GetLogger().Info("GetInventoryDetailUrl", zap.String("inventroyUrl", inventroyUrl), zap.String("name", name))

	priceArgs, invArgs := ezgo.Await2(
		ezgo.Async2(ezgo.Bind3_2(fetchJsonPathesWithRetry, httpClient, priceUrl, []string{"finalOnlinePrice", "discount"})),
		ezgo.Async2(ezgo.Bind3_2(fetchJsonPathesWithRetry, httpClient, inventroyUrl, []string{"invAvailable"})),
	)

	priceResult, err := priceArgs.Unpack()
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "fetchJsonPath.finalOnlinePrice."+priceUrl)
	}
	if len(priceResult) != 2 {
		return nil, fmt.Errorf("not enough price result count: %d", len(priceResult))
	}

	hasInvResult, err := invArgs.Unpack()
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "fetchJsonPath.invAvailable."+inventroyUrl)
	}
	if len(hasInvResult) != 1 {
		return nil, fmt.Errorf("not enough inventory result count: %d", len(hasInvResult))
	}

	return &ItemModel{
		ItemId:     itemId,
		CategoryId: categoryId,
		ProductId:  productId,
		Price:      priceResult[0].Float() - priceResult[1].Float(),
		Available:  hasInvResult[0].Bool(),
	}, nil
}

func FetchItemModelSequential(
	scope *ezgo.Scope,
	httpClient *ezgo.HttpClient,
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
	priceResult, err := fetchJsonPathesWithRetry(httpClient, priceUrl, []string{"finalOnlinePrice", "discount"})
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "fetchJsonPath.finalOnlinePrice."+priceUrl)
	}
	if len(priceResult) != 2 {
		return nil, ezgo.NewCausef(err, "not enough price result count: %d", len(priceResult))
	}

	// fetch inventory
	inventroyUrl := builGetInventoryDetailUrl(itemId, categoryId, productId, queryStringPatch)
	scope.GetLogger().Info("GetInventoryDetailUrl", zap.String("inventroyUrl", inventroyUrl), zap.String("name", name))
	hasInvResult, err := fetchJsonPathesWithRetry(httpClient, inventroyUrl, []string{"invAvailable"})
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "fetchJsonPath.invAvailable."+inventroyUrl)
	}

	return &ItemModel{
		ItemId:     itemId,
		CategoryId: categoryId,
		ProductId:  productId,
		Price:      priceResult[0].Float() - priceResult[1].Float(),
		Available:  hasInvResult[0].Bool(),
	}, nil
}

func fetchJsonPathesWithRetry(httpClient *ezgo.HttpClient, url string, pathes []string) ([]*gjson.Result, error) {
	return ezgo.RetryOnErr(
		ezgo.Bind3_2(fetchJsonPathes, httpClient, url, pathes),
		3,
		time.Second,
	)
}

func fetchJsonPathes(httpClient *ezgo.HttpClient, url string, pathes []string) ([]*gjson.Result, error) {
	body, err := httpClient.
		WithDefaultUserAgent().
		SetCookieString(getCookieString()).
		SetHeader("accept-language", "en-US,en;q=0.9").
		SetHeader("sec-ch-ua", `Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"`).
		SetHeader("sec-ch-ua-mobile", "?0").
		SetHeader("sec-ch-ua-platform", `"macOS"`).
		SetHeader("sec-fetch-dest", "document").
		Get(url)
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCausef(err, "HttpCall(%s)", url)
	}

	results := make([]*gjson.Result, len(pathes))
	for i, path := range pathes {
		results[i], err = ezgo.ExtractJsonPath(string(body), path)
		if ezgo.IsErr(err) {
			return nil, ezgo.NewCausef(err, "ExtractJsonPath(%s, %s)", ezgo.FirstNChars(body, 200), path)
		}
	}

	return results, nil
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
		//		WithQueryStringPatch(queryStringPatch). // somehow price api can't have the patch
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
