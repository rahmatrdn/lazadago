package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetMultipleOrderItemsOrderItemsResponseEntity struct{
    Reason	string	`json:"reason"`
    DigitalDeliveryInfo	string	`json:"digital_delivery_info"`
    PromisedShippingTime	string	`json:"promised_shipping_time"`
    OrderId	int	`json:"order_id"`
    VoucherAmount	string	`json:"voucher_amount"`
    ReturnStatus	string	`json:"return_status"`
    ShippingType	string	`json:"shipping_type"`
    ShipmentProvider	string	`json:"shipment_provider"`
    CancelReturnInitiator	string	`json:"cancel_return_initiator"`
    Variation	string	`json:"variation"`
    CreatedAt	string	`json:"created_at"`
    InvoiceNumber	string	`json:"invoice_number"`
    ShippingAmount	string	`json:"shipping_amount"`
    Currency	string	`json:"currency"`
    ShopId	string	`json:"shop_id"`
    Sku	string	`json:"sku"`
    VoucherCode	string	`json:"voucher_code"`
    WalletCredits	string	`json:"wallet_credits"`
    UpdatedAt	string	`json:"updated_at"`
    IsDigital	int	`json:"is_digital"`
    TrackingCodePre	string	`json:"tracking_code_pre"`
    OrderItemId	int	`json:"order_item_id"`
    PackageId	string	`json:"package_id"`
    TrackingCode	string	`json:"tracking_code"`
    ShippingServiceCost	int	`json:"shipping_service_cost"`
    ExtraAttributes	string	`json:"extra_attributes"`
    PaidPrice	string	`json:"paid_price"`
    ShippingProviderType	string	`json:"shipping_provider_type"`
    ProductDetailUrl	string	`json:"product_detail_url"`
    ShopSku	string	`json:"shop_sku"`
    ReasonDetail	string	`json:"reason_detail"`
    PurchaseOrderId	string	`json:"purchase_order_id"`
    PurchaseOrderNumber	string	`json:"purchase_order_number"`
    Name	string	`json:"name"`
    ProductMainImage	string	`json:"product_main_image"`
    ItemPrice	string	`json:"item_price"`
    TaxAmount	string	`json:"tax_amount"`
    Status	string	`json:"status"`
    VoucherPlatform	string	`json:"voucher_platform"`
    VoucherSeller	string	`json:"voucher_seller"`
    OrderType	string	`json:"order_type"`
    StagePayStatus	string	`json:"stage_pay_status"`
    OrderFlag	string	`json:"order_flag"`
    SlaTimeStamp	string	`json:"sla_time_stamp"`
    WarehouseCode	string	`json:"warehouse_code"`
    ShippingFeeOriginal	string	`json:"shipping_fee_original"`
    ShippingFeeDiscountSeller	string	`json:"shipping_fee_discount_seller"`
    ShippingFeeDiscountPlatform	string	`json:"shipping_fee_discount_platform"`
    VoucherCodeSeller	string	`json:"voucher_code_seller"`
    VoucherCodePlatform	string	`json:"voucher_code_platform"`
    DeliveryOptionSof	string	`json:"delivery_option_sof"`
    IsFbl	string	`json:"is_fbl"`
    IsReroute	string	`json:"is_reroute"`
    VoucherSellerLpi	string	`json:"voucher_seller_lpi"`
    VoucherPlatformLpi	string	`json:"voucher_platform_lpi"`
    BuyerId	string	`json:"buyer_id"`
}
func (g GetMultipleOrderItemsOrderItemsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
