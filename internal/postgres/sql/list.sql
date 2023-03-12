SELECT
    Id,
    MerchantId,
    CardHolderName,
    CardNumber,
    ExpiryMonth,
    ExpiryYear,
    Amount,
    CurrencyCode,
    Reference,
    Status,
    CreatedAt,
    UpdatedAt
FROM Payment
WHERE MerchantId = @MerchantId