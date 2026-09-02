package domain

import "errors"

var (
	// ErrWalletNotFound is returned when a transfer references a wallet
	// that does not exist.
	ErrWalletNotFound = errors.New("wallet not found")

	// ErrInsufficientFunds is returned when the source wallet does not
	// have enough balance to cover the transfer.
	ErrInsufficientFunds = errors.New("insufficient funds")

	// ErrIdempotencyConflict is returned when an idempotency key is reused
	// with a request payload that differs from the original.
	ErrIdempotencyConflict = errors.New("idempotency key reused with a different request payload")

	// ErrProcessingTimeout is returned when a duplicate request waited too
	// long for the original in-flight request to complete. The caller may
	// safely retry: the original transaction either commits (and a retry
	// will replay it) or is rolled back (and a retry reprocesses cleanly).
	ErrProcessingTimeout = errors.New("original request is still processing, retry")

	// ErrInvalidAmount is returned when the transfer amount is not positive.
	ErrInvalidAmount = errors.New("amount must be a positive integer")

	// ErrSameWallet is returned when fromWalletId and toWalletId are equal.
	ErrSameWallet = errors.New("fromWalletId and toWalletId must differ")

	// ErrMissingField is returned when a required request field is empty.
	ErrMissingField = errors.New("missing required field")
)
