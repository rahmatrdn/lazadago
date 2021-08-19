package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrdersOrdersResponseEntity struct{
    BranchNumber	string	`json:"branch_number"`
    TaxCode	string	`json:"tax_code"`
    ExtraAttributes	string	`json:"extra_attributes"`
    AddressUpdatedAt	string	`json:"address_updated_at"`
    ShippingFee	string	`json:"shipping_fee"`
    CustomerFirstName	string	`json:"customer_first_name"`
    PaymentMethod	string	`json:"payment_method"`
    Statuses	[]string	`json:"statuses"`
    Remarks	string	`json:"remarks"`
    OrderNumber	string	`json:"order_number"`
    OrderId	string	`json:"order_id"`
    Voucher	string	`json:"voucher"`
    NationalRegistrationNumber	string	`json:"national_registration_number"`
    PromisedShippingTimes	string	`json:"promised_shipping_times"`
    ItemsCount	int	`json:"items_count"`
    VoucherPlatform	string	`json:"voucher_platform"`
    VoucherSeller	string	`json:"voucher_seller"`
    CreatedAt	string	`json:"created_at"`
    Price	string	`json:"price"`
    AddressBilling	GetOrdersAddressBillingResponseEntity	`json:"address_billing"`
    WarehouseCode	string	`json:"warehouse_code"`
    ShippingFeeOriginal	string	`json:"shipping_fee_original"`
    ShippingFeeDiscountSeller	string	`json:"shipping_fee_discount_seller"`
    ShippingFeeDiscountPlatform	string	`json:"shipping_fee_discount_platform"`
    AddressShipping	GetOrdersAddressShippingResponseEntity	`json:"address_shipping"`
    CustomerLastName	string	`json:"customer_last_name"`
    GiftOption	string	`json:"gift_option"`
    VoucherCode	string	`json:"voucher_code"`
    UpdatedAt	string	`json:"updated_at"`
    DeliveryInfo	string	`json:"delivery_info"`
    GiftMessage	string	`json:"gift_message"`
}
func (g GetOrdersOrdersResponseEntity) String() string {
    return lib.ObjectToString(g)
}
