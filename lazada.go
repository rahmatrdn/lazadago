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
    "github.com/wjpxxx/lazadago/finance"
	financeentity "github.com/wjpxxx/lazadago/finance/entity"
    "github.com/wjpxxx/lazadago/logistics"
	logisticsentity "github.com/wjpxxx/lazadago/logistics/entity"
    "github.com/wjpxxx/lazadago/seller"
	sellerentity "github.com/wjpxxx/lazadago/seller/entity"
    "github.com/wjpxxx/lazadago/datamoat"
	datamoatentity "github.com/wjpxxx/lazadago/datamoat/entity"
    "github.com/wjpxxx/lazadago/fbl"
	fblentity "github.com/wjpxxx/lazadago/fbl/entity"
    "github.com/wjpxxx/lazadago/gspproduct"
	gspproductentity "github.com/wjpxxx/lazadago/gspproduct/entity"
    "github.com/wjpxxx/lazadago/etickets"
	eticketsentity "github.com/wjpxxx/lazadago/etickets/entity"
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
	//Finance
    GetPayoutStatus (createdAfter string) financeentity.GetPayoutStatusResult
    GetTransactionDetails (transType string,startTime string,endTime string,limit string,offset string) financeentity.GetTransactionDetailsResult
    QueryTransactionDetails (offset string,transType string,tradeOrderId string,limit string,startTime string,endTime string,tradeOrderLineId string) financeentity.QueryTransactionDetailsResult
    //Logistics
    GetOrderTrace (sellerId string,orderId string,locale string,ofcPackageIdList []string) logisticsentity.GetOrderTraceResult
    GetShipmentProviders () logisticsentity.GetShipmentProvidersResult
    //Seller
    GetMultiWarehouseBySeller (addressTypes []string) sellerentity.GetMultiWarehouseBySellerResult
    GetSeller () sellerentity.GetSellerResult
    GetSellerMetricsById () sellerentity.GetSellerMetricsByIdResult
    GetSellerPerformance (language string) sellerentity.GetSellerPerformanceResult
    UpdateSeller (payload string) sellerentity.UpdateSellerResult
    UpdateUser (payload string) sellerentity.UpdateUserResult
    //DataMoat
    DataMoatComputeRisk (time string,appName string,userId string,userIp string,ati string) datamoatentity.DataMoatComputeRiskResult
    DataMoatLogin (time string,appName string,userId string,tid string,userIp string,ati string,loginResult string,loginMessage string) datamoatentity.DataMoatLoginResult
    //FBL
    GetPlatformProducts (perPage int,sellerId int,marketplace string,sellerSku string,platformSkuName string,readyForInbound bool,platformSku string,page int) fblentity.GetPlatformProductsResult
    GetFulfillmentProductDetail (perPage int,shelfLifeFlag bool,marketplace string,fulfillmentSku string,serialNumberFlag bool,page int,fulfillmentSkuName string,barcode string) fblentity.GetFulfillmentProductDetailResult
    GetInboundOrderDetail (inboundOrderNo string,marketplace string) fblentity.GetInboundOrderDetailResult
    GetInboundOrderList (inboundOrderNo string,creationTimeFrom string,creationTimeTo string,inboundWarehouse string,sellerSku string,fulfillmentSku string,marketplace string,page string,perPage string) fblentity.GetInboundOrderListResult
    GetInventoryChangedSKU (warehouseCode string,page int,perPage int,marketPlace string,operateTimeFrom string,operateTimeTo string) fblentity.GetInventoryChangedSKUResult
    GetInventoryOperateLog (page int,perPage int,marketPlace string,operateTimeFrom string,operateTimeTo string,warehouseCode string,fulfillmentSkuId string) fblentity.GetInventoryOperateLogResult
    GetOutboundOrderDetail (outboundOrderNo string,marketplace string) fblentity.GetOutboundOrderDetailResult
    GetOutboundOrderList (outboundOrderNo string,creationTimeFrom string,creationTimeTo string,outboundWarehouse string,sellerSku string,fulfillmentSku string,marketplace string,page string,perPage string) fblentity.GetOutboundOrderListResult
    GetWarehouseStock (sellerSku string,marketplace string,fulfilmentSku string,storeCode string) fblentity.GetWarehouseStockResult
    GetWarehouseStockV3 (sellerSku string,marketplace string,fulfilmentSku string,storeCode string) fblentity.GetWarehouseStockV3Result
    UploadWaybill (waybill []byte,packageCode string,trackingNumber string,extendsField string,storeCode string) fblentity.UploadWaybillResult
    //GSPProduct API
    CreateGlobalProduct (payload string) gspproductentity.CreateGlobalProductResult
    GetGlobalProductStatus (params gspproductentity.GetGlobalProductStatusParamsRequestEntity) gspproductentity.GetGlobalProductStatusResult
    GetUnfilledAttribute (offset int,limit int,attributeTag string) gspproductentity.GetUnfilledAttributeResult
    UpdateGlobalProductAttribute (payload string) gspproductentity.UpdateGlobalProductAttributeResult
    //E-Tickets API
    GetOrderItemsFromBarCode (code string) eticketsentity.GetOrderItemsFromBarCodeResult
    RedeemOrderItems (bizType int,code string,outerId string,serialNum string,consumeNum int,storeId string,posId string) eticketsentity.RedeemOrderItemsResult
    GlobalEticketMerchantMaAvailable (bizType int,code string,serialNum string,posId string,outerId string,consumeNum int,consumeStoreId string) eticketsentity.GlobalEticketMerchantMaAvailableResult
    GlobalEticketMerchantMaConsume (bizType int,serialNum string,posId string,outerId string,consumeNum int,code string,consumeStoreId string) eticketsentity.GlobalEticketMerchantMaConsumeResult
    GlobalEticketMerchantMaFailsend (bizType int,subCode string,outerId string,subMsg string) eticketsentity.GlobalEticketMerchantMaFailsendResult
    GlobalEticketMerchantMaQuery (code string,sellerId int,storeId int) eticketsentity.GlobalEticketMerchantMaQueryResult
    GlobalEticketMerchantMaQueryTbMa (code string) eticketsentity.GlobalEticketMerchantMaQueryTbMaResult
    GlobalEticketMerchantMaSend (bizType int,isvMaList []eticketsentity.GlobalEticketMerchantMaSendIsvMaListRequestEntity,outerId string) eticketsentity.GlobalEticketMerchantMaSendResult

}

//Lazada
type Lazada struct {
	auth.Auth
	system.System
	order.Order
	product.Product
    finance.Finance
    logistics.Logistics
    seller.Seller
    datamoat.DataMoat
    fbl.Fbl
    gspproduct.GspProduct
    etickets.ETickets
}
//SetAccessToken 设置token
func (l *Lazada)SetAccessToken(accessToken string){
    l.Auth.Config.SetAccessToken(accessToken)
    l.System.Config.SetAccessToken(accessToken)
    l.Order.Config.SetAccessToken(accessToken)
    l.Product.Config.SetAccessToken(accessToken)
    l.Finance.Config.SetAccessToken(accessToken)
    l.Logistics.Config.SetAccessToken(accessToken)
    l.Seller.Config.SetAccessToken(accessToken)
    l.DataMoat.Config.SetAccessToken(accessToken)
    l.Fbl.Config.SetAccessToken(accessToken)
    l.GspProduct.Config.SetAccessToken(accessToken)
    l.ETickets.Config.SetAccessToken(accessToken)
}

//NewApi
func NewApi(cfg *lazadaConfig.Config)Lazadar{
	return &Lazada{
		auth.Auth{Config: cfg},
		system.System{Config: cfg},
		order.Order{Config: cfg},
		product.Product{Config: cfg},
        finance.Finance{Config:cfg},
        logistics.Logistics{Config: cfg},
        seller.Seller{Config: cfg},
        datamoat.DataMoat{Config: cfg},
        fbl.Fbl{Config: cfg},
        gspproduct.GspProduct{Config: cfg},
        etickets.ETickets{Config: cfg},
	}
}
