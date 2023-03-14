package internal

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ref: https://rotational.io/blog/marshaling-go-enums-to-and-from-json/
type PaymentStatus int

const (
	PaymentStatusPending PaymentStatus = iota + 1
	PaymentStatusApproved
	PaymentStatusDeclined
)

const (
	pendingName  = "pending"
	approvedName = "approved"
	declinedName = "declined"
)

var (
	PaymentStatusName = map[PaymentStatus]string{
		PaymentStatusApproved: approvedName,
		PaymentStatusDeclined: declinedName,
	}

	PaymentStatusValue = map[string]PaymentStatus{
		approvedName: PaymentStatusApproved,
		declinedName: PaymentStatusDeclined,
	}
)

func (p PaymentStatus) String() string {
	return PaymentStatusName[p]
}

func ParsePaymentStatus(s string) (PaymentStatus, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := PaymentStatusValue[s]
	if !ok {
		return PaymentStatus(0), fmt.Errorf("%s is not a valid payment status", s)
	}

	return value, nil
}

func (p PaymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PaymentStatus) UnmarshalJSON(b []byte) (err error) {
	var paymentStatusName string
	if err := json.Unmarshal(b, &paymentStatusName); err != nil {
		return err
	}
	if *p, err = ParsePaymentStatus(paymentStatusName); err != nil {
		return err
	}

	return nil
}
