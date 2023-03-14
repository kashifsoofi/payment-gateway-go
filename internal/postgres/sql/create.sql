INSERT INTO payments (
    id,
    merchant_id,
    card_holder_name,
    card_number,
    expiry_month,
    expiry_year,
    amount,
    currency_code,
    reference,
    status,
    created_at,
    updated_at
)
VALUES (
    @id,
    @merchant_id,
    @card_holder_name,
    @card_number,
    @expiry_month,
    @expiry_year,
    @amount,
    @currency_code,
    @reference,
    @status,
    @created_at,
    @updated_at
)