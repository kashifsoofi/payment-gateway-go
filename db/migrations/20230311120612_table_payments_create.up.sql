CREATE TABLE IF NOT EXISTS payments (
    clustered_id        BIGSERIAL,
    id                  uuid            NOT NULL,
    merchant_id         uuid            NOT NULL,
    card_holder_name    VARCHAR(255)    NOT NULL,
    card_number         CHAR(4)         NOT NULL,
    expiry_month        INT             NOT NULL,
    expiry_year         INT             NOT NULL,
    amount              DECIMAL(12, 4)  NOT NULL,
    currency_code       CHAR(3)         NOT NULL,
    reference           VARCHAR(50)     NOT NULL,
    status              VARCHAR(40)     NOT NULL,
    created_at          TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    updated_at          TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    PRIMARY KEY (clustered_id),
    UNIQUE (id)
);