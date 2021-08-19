package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrderDataResponseEntity struct{
    AddressShipping	GetOrderAddressShippingResponseEntity	`json:"address_shipping"`
    CustomerLastName	string	`json:"customer_last_name"`
    GiftOption	bool	`json:"gift_option"`
    VoucherCode	string	`json:"voucher_code"`
    UpdatedAt	string	`json:"updated_at"`
    DeliveryInfo	string	`json:"delivery_info"`
    GiftMessage	string	`json:"gift_message"`
    BranchNumber	string	`json:"branch_number"`
    TaxCode	string	`json:"tax_code"`
    ExtraAttributes	string	`json:"extra_attributes"`
    ShippingFee	string	`json:"shipping_fee"`
    CustomerFirstName	string	`json:"customer_first_name"`
    PaymentMethod	string	`json:"payment_method"`
    Statuses	[]string	`json:"statuses"`
    Remarks	string	`json:"remarks"`
    OrderNumber	int	`json:"order_number"`
    OrderId	int	`json:"order_id"`
    Voucher	string	`json:"voucher"`
    NationalRegistrationNumber	string	`json:"national_registration_number"`
    PromisedShippingTimes	string	`json:"promised_shipping_times"`
    ItemsCount	int	`json:"items_count"`
    CreatedAt	string	`json:"created_at"`
    Price	string	`json:"price"`
    AddressBilling	GetOrderAddressBillingResponseEntity	`json:"address_billing"`
    WarehouseCode	string	`json:"warehouse_code"`
    ShippingFeeOriginal	string	`json:"shipping_fee_original"`
    ShippingFeeDiscountSeller	string	`json:"shipping_fee_discount_seller"`
    ShippingFeeDiscountPlatform	string	`json:"shipping_fee_discount_platform"`
}
func (g GetOrderDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
