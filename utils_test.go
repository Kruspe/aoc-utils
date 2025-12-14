package aoc_utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type utilSuite struct {
	suite.Suite
}

func Test_UtilSuite(t *testing.T) {
	suite.Run(t, &utilSuite{})
}

func (s *utilSuite) Test_ReadGCD() {
	gcd := GCD(10, 5)
	require.Equal(s.T(), 5, gcd)
	gcd = GCD(2000, 45)
	require.Equal(s.T(), 5, gcd)
}

func (s *utilSuite) Test_LCM() {
	lcm, err := LCM([]int{2000, 45, 400, 234})
	require.NoError(s.T(), err)
	require.Equal(s.T(), 234000, lcm)
	lcm, err = LCM([]int{123, 456, 234})
	require.NoError(s.T(), err)
	require.Equal(s.T(), 729144, lcm)
}

func (s *utilSuite) Test_Permutations() {
	combs := Combinations([]any{1, 2, 3, 4})
	s.Len(combs, 15)
}
