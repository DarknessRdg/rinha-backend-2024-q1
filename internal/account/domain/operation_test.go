package domain

import (
	"testing"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
	"github.com/stretchr/testify/require"
)

func TestOperationFromString(t *testing.T) {
	tests := []struct {
		name              string
		operation         string
		expectedOperation Operation
		expectError       bool
	}{
		{
			name:              "Credit operation",
			operation:         "c",
			expectedOperation: OperationCredit,
		},
		{
			name:              "Debit operation",
			operation:         "d",
			expectedOperation: OperationDebit,
		},
		{
			name:        "Unknown operation",
			operation:   "a",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			op, err := OperationFromString(test.operation)

			if test.expectError {
				require.Error(t, err)
				require.ErrorIs(t, err, errs.UnprocessableEntity(""))
				require.Equal(t, -1, int(op))
				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expectedOperation, op)
		})
	}
}
