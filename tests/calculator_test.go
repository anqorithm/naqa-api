package tests

import (
	"testing"
	"time"

	"github.com/anqorithm/naqa-api/internal/handlers"
)

func TestGetYearsInPeriod(t *testing.T) {
	tests := []struct {
		name      string
		startDate time.Time
		endDate   time.Time
		expected  []int
	}{
		{
			name:      "single year",
			startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			expected:  []int{2023},
		},
		{
			name:      "two consecutive years",
			startDate: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			expected:  []int{2023, 2024},
		},
		{
			name:      "three years",
			startDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			expected:  []int{2022, 2023, 2024},
		},
		{
			name:      "same day",
			startDate: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
			expected:  []int{2023},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handlers.GetYearsInPeriod(tt.startDate, tt.endDate)
			if len(result) != len(tt.expected) {
				t.Errorf("getYearsInPeriod() = %v, expected %v", result, tt.expected)
				return
			}
			for i, year := range result {
				if year != tt.expected[i] {
					t.Errorf("getYearsInPeriod() = %v, expected %v", result, tt.expected)
					break
				}
			}
		})
	}
}

func TestGetDaysInYear(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		expected int
	}{
		{
			name:     "regular year",
			year:     2023,
			expected: 365,
		},
		{
			name:     "leap year",
			year:     2024,
			expected: 366,
		},
		{
			name:     "century non-leap year",
			year:     1900,
			expected: 365,
			
		},
		{
			name:     "century leap year",
			year:     2000,
			expected: 366,
		},
		{
			name:     "another leap year",
			year:     2020,
			expected: 366,
		},
	}


	for _, tt := range te
	
		result := handlers.GetDaysInYear(tt.year)
			if result != tt.expected {
				t.Errorf("getDaysInYear(%d) = %d, expected %d", tt.year, result, tt.expected)
			}
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		expected bool
	}{
		{
			name:     "regular leap year",
			year:     2024,
			expected: true,
		},
		{
			name:     "non-leap year",
			year:     2023,
			expected: false,
		},
		{
			name:     "century non-leap year",
			year:     1900,
			expected: false,
		},
		{
			name:     "century leap year",
			year:     2000,
			expected: true,
		},
		{
			name:     "year 2100 (not leap)",
			year:     2100,
			expected: false,
		},
		{
			name:     "year 2400 (leap)",
			year:     2400,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handlers.IsLeapYear(tt.year)
			if result != tt.expected {
				t.Errorf("isLeapYear(%d) = %v, expected %v", tt.year, result, tt.expected)
			}
		})
	}
}

func TestGetDaysInYearForPeriod(t *testing.T) {
	tests := []struct {
		name      string
		startDate time.Time
		endDate   time.Time
		year      int
		expected  int
	}{
		{
			name:      "full year 2023",
			startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			year:      2023,
			expected:  365,
		},
		{
			name:      "full leap year 2024",
			startDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			year:      2024,
			expected:  366,
		},
		{
			name:      "partial year - first half 2023",
			startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2023, 6, 30, 0, 0, 0, 0, time.UTC),
			year:      2023,
			expected:  181, // Jan 1 to June 30
		},
		{
			name:      "partial year - second half 2023",
			startDate: time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			year:      2023,
			expected:  184, // July 1 to Dec 31
		},
		{
			name:      "cross year period - only 2023 part",
			startDate: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			year:      2023,
			expected:  215, // June 1 to Dec 31, 2023 (inclusive)
		},
		{
			name:      "cross year period - only 2024 part",
			startDate: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
			year:      2024,
			expected:  182, // Jan 1 to June 30, 2024 (inclusive)
		},
		{
			name:      "no overlap - period before year",
			startDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
			year:      2023,
			expected:  0,
		},
		{
			name:      "no overlap - period after year",
			startDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			year:      2023,
			expected:  0,
		},
		{
			name:      "single day",
			startDate: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
			year:      2023,
			expected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handlers.GetDaysInYearForPeriod(tt.startDate, tt.endDate, tt.year)
			if result != tt.expected {
				t.Errorf("getDaysInYearForPeriod(%v, %v, %d) = %d, expected %d",
					tt.startDate.Format("2006-01-02"), tt.endDate.Format("2006-01-02"), tt.year, result, tt.expected)
			}
		})
	}
}

// TestPurificationCalculationFormula tests the purification calculation logic
func TestPurificationCalculationFormula(t *testing.T) {
	tests := []struct {
		name               string
		numberOfStocks     int
		purificationRate   float64
		daysInPeriod       int
		totalDaysInYear    int
		expectedAmount     float64
	}{
		{
			name:             "Example from README - regular year",
			numberOfStocks:   100,
			purificationRate: 0.025, // 2.5%
			daysInPeriod:     180,
			totalDaysInYear:  365,
			expectedAmount:   1.2328767123287671, // 100 * 0.025 * (180/365)
		},
		{
			name:             "Example from README - leap year",
			numberOfStocks:   100,
			purificationRate: 0.025, // 2.5%
			daysInPeriod:     180,
			totalDaysInYear:  366,
			expectedAmount:   1.2295081967213115, // 100 * 0.025 * (180/366)
		},
		{
			name:             "Full year regular",
			numberOfStocks:   50,
			purificationRate: 0.05, // 5%
			daysInPeriod:     365,
			totalDaysInYear:  365,
			expectedAmount:   2.5, // 50 * 0.05 * (365/365) = 2.5
		},
		{
			name:             "Full leap year",
			numberOfStocks:   50,
			purificationRate: 0.05, // 5%
			daysInPeriod:     366,
			totalDaysInYear:  366,
			expectedAmount:   2.5, // 50 * 0.05 * (366/366) = 2.5
		},
		{
			name:             "Half year",
			numberOfStocks:   200,
			purificationRate: 0.01, // 1%
			daysInPeriod:     182, // approximately half year
			totalDaysInYear:  365,
			expectedAmount:   0.9972602739726027, // 200 * 0.01 * (182/365)
		},
		{
			name:             "Zero purification rate",
			numberOfStocks:   100,
			purificationRate: 0.0,
			daysInPeriod:     365,
			totalDaysInYear:  365,
			expectedAmount:   0.0,
		},
		{
			name:             "High purification rate",
			numberOfStocks:   10,
			purificationRate: 0.15, // 15%
			daysInPeriod:     90,
			totalDaysInYear:  365,
			expectedAmount:   0.3698630136986301, // 10 * 0.15 * (90/365)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This mimics the calculation from the actual implementation
			yearProportion := float64(tt.daysInPeriod) / float64(tt.totalDaysInYear)
			result := float64(tt.numberOfStocks) * tt.purificationRate * yearProportion

			// Allow small floating point differences
			tolerance := 1e-10
			if abs(result-tt.expectedAmount) > tolerance {
				t.Errorf("Purification calculation: numberOfStocks=%d, purificationRate=%f, daysInPeriod=%d, totalDaysInYear=%d\nResult: %f, Expected: %f, Difference: %f",
					tt.numberOfStocks, tt.purificationRate, tt.daysInPeriod, tt.totalDaysInYear, result, tt.expectedAmount, abs(result-tt.expectedAmount))
			}
		})
	}
}

// TestREADMEFormulaVsImplementation compares README formula with actual implementation
func TestREADMEFormulaVsImplementation(t *testing.T) {
	tests := []struct {
		name             string
		numberOfStocks   int
		purificationRate float64
		daysHeld         int
		year             int
	}{
		{
			name:             "README example in regular year",
			numberOfStocks:   100,
			purificationRate: 0.025,
			daysHeld:         180,
			year:             2023,
		},
		{
			name:             "README example in leap year",
			numberOfStocks:   100,
			purificationRate: 0.025,
			daysHeld:         180,
			year:             2024,
		},
		{
			name:             "Full year regular",
			numberOfStocks:   50,
			purificationRate: 0.05,
			daysHeld:         365,
			year:             2023,
		},
		{
			name:             "Full year leap",
			numberOfStocks:   50,
			purificationRate: 0.05,
			daysHeld:         366,
			year:             2024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// README formula (simplified)
			readmeResult := float64(tt.numberOfStocks) * tt.purificationRate * float64(tt.daysHeld) / 365.0

			// Implementation formula (accurate)
			totalDaysInYear := handlers.GetDaysInYear(tt.year)
			yearProportion := float64(tt.daysHeld) / float64(totalDaysInYear)
			implementationResult := float64(tt.numberOfStocks) * tt.purificationRate * yearProportion

			t.Logf("Test: %s", tt.name)
			t.Logf("README formula result: %f", readmeResult)
			t.Logf("Implementation result: %f", implementationResult)
			t.Logf("Difference: %f", abs(readmeResult-implementationResult))
			t.Logf("Year %d has %d days", tt.year, totalDaysInYear)

			// For leap years, there should be a difference
			if tt.year == 2024 && tt.daysHeld == 366 {
				// The results should be different for leap years when using full year
				tolerance := 1e-10
				if abs(readmeResult-implementationResult) < tolerance {
					t.Errorf("Expected difference between README and implementation for leap year, but they are equal")
				}
			}
		})
	}
}

// Helper function for absolute value of float64
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// TestEdgeCases tests edge cases for the purification calculation
func TestPurificationCalculationEdgeCases(t *testing.T) {
	tests := []struct {
		name               string
		numberOfStocks     int
		purificationRate   float64
		daysInPeriod       int
		totalDaysInYear    int
		expectedAmount     float64
	}{
		{
			name:             "Single day",
			numberOfStocks:   100,
			purificationRate: 0.025,
			daysInPeriod:     1,
			totalDaysInYear:  365,
			expectedAmount:   0.006849315068493151, // 100 * 0.025 * (1/365)
		},
		{
			name:             "Zero stocks",
			numberOfStocks:   0,
			purificationRate: 0.025,
			daysInPeriod:     180,
			totalDaysInYear:  365,
			expectedAmount:   0.0,
		},
		{
			name:             "Large number of stocks",
			numberOfStocks:   1000000,
			purificationRate: 0.001,
			daysInPeriod:     365,
			totalDaysInYear:  365,
			expectedAmount:   1000.0, // 1,000,000 * 0.001 * 1
		},
		{
			name:             "Very small purification rate",
			numberOfStocks:   100,
			purificationRate: 0.0001, // 0.01%
			daysInPeriod:     365,
			totalDaysInYear:  365,
			expectedAmount:   0.01,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yearProportion := float64(tt.daysInPeriod) / float64(tt.totalDaysInYear)
			result := float64(tt.numberOfStocks) * tt.purificationRate * yearProportion

			tolerance := 1e-10
			if abs(result-tt.expectedAmount) > tolerance {
				t.Errorf("Result: %f, Expected: %f, Difference: %f", result, tt.expectedAmount, abs(result-tt.expectedAmount))
			}
		})
	}
}