package domain

import "github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"

type Operation int

const (
	OperationCredit Operation = iota
	OperationDebit  Operation = iota
)

func OperationFromString(operation string) (Operation, error) {
	switch operation {
	case "c":
		return OperationCredit, nil
	case "d":
		return OperationDebit, nil
	}

	return -1, errs.UnprocessableEntity("invalid operation type")
}
