## v1.6.0
- **Feature:** Add new `Assets` model for managing service certificate assets
- **Feature:** Add new `LocalizedVersion` model for localized content management
- **Feature:** Add new `NoticePeriod` model with types: `SAME_DAY`, `DAYS`, `MONTHS`
- **Feature:** Add new `ServiceCertificate` model for service certification
- **Feature:** Add `Assets` field to `CatalogProductDetail` model
- **Feature:** Add `OfferType` field to `CatalogProductDetail` model
- **Feature:** Add `NoticePeriod` field to `CatalogProductPricingOption` model
- Add `required:"true"` tags to model structs

## v1.5.0 (2025-06-10)
- **Feature:** Add new field `Facets` in `ListCatalogProductsResponse`

## v1.4.0 (2025-06-06)
- **Fix:** Fixed types for `VendorId`, `ProjectId`, `OrganizationId` and `SubscriptionId`

## v1.3.1 (2025-06-04)
- **Feature:** Added `Industries` attribute

## v1.3.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v1.2.1 (2025-05-15)
- **Feature:** Add new method `VendorsSubscriptionsReject`

## v1.2.0 (2025-05-14)
- **Breaking change:** Introduce typed enum constants for status attributes

## v1.1.0 (2025-05-13)
- **Breaking Change:** Added organization id to `VendorSubscription`

## v1.0.1 (2025-05-09)
- **Feature:** Update user-agent header

## v1.0.0 (2025-05-02)
- **Breaking Change:**
    - Introduced dedicated type for product id with appropriate validations
- **Feature:** 
    - Improved nil-safety
    - subscription products contain the plan id

## v0.5.0 (2025-04-11)
- **Feature:** Add new `InquiryContactSales`, `InquirySuggestProduct`, `PriceType`, `PricingOption` and `DeliveryMethod`

## v0.4.0 (2025-04-04)
- **Feature:** Add new `VendorProductId` attribute for subscription products

## v0.3.1 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v0.3.0 (2025-02-25)
- **Feature:** Add method to create inquiries: `InquiriesCreateInquiry`
- **Feature:** Add `sort` property to `ApiListCatalogProductsRequest`
- **Feature:** Add payload `ApproveSubscriptionPayload` for `ApiApproveSubscriptionRequest`

## v0.2.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.1.0 (2025-01-10)

- **New**: STACKIT Marketplace module can be used to manage the STACKIT Marketplace.
