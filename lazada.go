package lazadago

import (
	"github.com/wjpxxx/lazadago/auth"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	"github.com/wjpxxx/lazadago/system"
	systementity "github.com/wjpxxx/lazadago/system/entity"
	"github.com/wjpxxx/lazadago/order"
	orderentity "github.com/wjpxxx/lazadago/order/entity"
	"github.com/wjpxxx/lazadago/product"
	productentity "github.com/wjpxxx/lazadago/product/entity"
)

//Lazadar
type Lazadar interface {
    SetAccessToken(accessToken string)
	//auth
	AuthorizationURL(redirectUri string) string
	//system
	GenerateAccessToken (code string,uuid string) systementity.GenerateAccessTokenResult
	RefreshAccessToken (refreshToken string) systementity.RefreshAccessTokenResult
	//order
	GetAwbDocumentHtml (orderItemIds string) orderentity.GetAwbDocumentHtmlResult
    GetAwbDocumentPDF (orderItemIds string) orderentity.GetAwbDocumentPDFResult     
    GetDocument (docType string,orderItemIds string) orderentity.GetDocumentResult     
    GetFailureReasons () orderentity.GetFailureReasonsResult     
    GetMultipleOrderItems (orderIds []int) orderentity.GetMultipleOrderItemsResult     
    GetOVOOrders (tradeOrderIds string) orderentity.GetOVOOrdersResult     
    GetOrder (orderId int) orderentity.GetOrderResult     
    GetOrderItems (orderId int) orderentity.GetOrderItemsResult     
    GetOrders (updateBefore string,sortDirection string,offset int,limit int,updateAfter string,sortBy string,createdBefore string,createdAfter string,status string) orderentity.GetOrdersResult     
    SetInvoiceNumber (orderItemId int,invoiceNumber string) orderentity.SetInvoiceNumberResult     
    SetRepack (packageId string) orderentity.SetRepackResult     
    SetStatusToCanceled (reasonDetail string,reasonId int,orderItemId int) orderentity.SetStatusToCanceledResult     
    SetStatusToPackedByMarketplace (shippingProvider string,deliveryType string,orderItemIds string) orderentity.SetStatusToPackedByMarketplaceResult     
    SetStatusToReadyToShip (deliveryType string,orderItemIds string,shipmentProvider string,trackingNumber string) orderentity.SetStatusToReadyToShipResult     
    SetStatusToSOFDelivered (orderItemIds string) orderentity.SetStatusToSOFDeliveredResult     
    SetStatusToSOFFailedDelivery (orderItemIds string) orderentity.SetStatusToSOFFailedDeliveryResult
	//product
    CreateProduct (payload string) productentity.CreateProductResult
    DeactivateProduct (apiRequestBody string) productentity.DeactivateProductResult
    GetBrandByPages (startRow string,pageSize string,languageCode string) productentity.GetBrandByPagesResult
    GetCategoryAttributes (primaryCategoryId string,languageCode string) productentity.GetCategoryAttributesResult
    GetCategorySuggestion (productName string) productentity.GetCategorySuggestionResult
    GetCategoryTree (languageCode string) productentity.GetCategoryTreeResult
    GetProductItem (itemId int,sellerSku string) productentity.GetProductItemResult
    GetProducts (filter string,updateBefore string,createBefore string,offset string,createAfter string,updateAfter string,limit string,options string,skuSellerList string) productentity.GetProductsResult
    GetQcStatus (offset int,limit int,sellerSkus []string) productentity.GetQcStatusResult
    GetResponse (batchId string) productentity.GetResponseResult
    GetSellerItemLimit () productentity.GetSellerItemLimitResult
    GetUnfilledAttributeItem (pageIndex int,attributeTag string,pageSize int,languageCode string) productentity.GetUnfilledAttributeItemResult
    MigrateImage (payload string) productentity.MigrateImageResult
    MigrateImages (payload string) productentity.MigrateImagesResult
    RemoveProduct (sellerSkuList string) productentity.RemoveProductResult
    RemoveSku (payload string) productentity.RemoveSkuResult
    SetImages (payload string) productentity.SetImagesResult
    UpdatePriceQuantity (payload string) productentity.UpdatePriceQuantityResult
    UpdateProduct (payload string) productentity.UpdateProductResult
    UploadImage (image []byte) productentity.UploadImageResult
    RetailFulfilmentCreate (platformName string,source string,sellerId int,platformSkuCode string,itemId int,skuId int,platformSkuName string,barcodeList []string,categoryId int,brand string,brandName string,isShelfLifeMgt bool,lifeCycleDays int,rejectLifeCycleDays int,lockupLifeCycleDays int,adventLifeCycleDays int,isSnMgt bool,cpWeight int,cpLength int,cpWidth int,cpHeight int,skuPrice int,features productentity.RetailFulfilmentCreateFeaturesRequestEntity) productentity.RetailFulfilmentCreateResult
    RetailFulfilmentUpdate (scItemId string,fulfillmentSkuName string,barcodeList []string,categoryId int,brand string,brandName string,isShelfLifeMgt bool,lifeCycleDays int,rejectLifeCycleDays int,lockupLifeCycleDays int,adventLifeCycleDays int,isSnMgt bool,cpWeight int,cpLength int,cpWidth int,cpHeight int,skuPrice int,features productentity.RetailFulfilmentUpdateFeaturesRequestEntity,source string) productentity.RetailFulfilmentUpdateResult
	//
}

//Lazada
type Lazada struct {
	auth.Auth
	system.System
	order.Order
	product.Product
}
//SetAccessToken 设置token
func (l *Lazada)SetAccessToken(accessToken string){
    l.Auth.Config.SetAccessToken(accessToken)
    l.System.Config.SetAccessToken(accessToken)
    l.Order.Config.SetAccessToken(accessToken)
    l.Product.Config.SetAccessToken(accessToken)
}

//NewApi
func NewApi(cfg *lazadaConfig.Config)Lazadar{
	return &Lazada{
		auth.Auth{Config: cfg},
		system.System{Config: cfg},
		order.Order{Config: cfg},
		product.Product{Config: cfg},
	}
}
