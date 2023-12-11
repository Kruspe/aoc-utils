package aoc_utils

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type utilSuite struct {
	suite.Suite
}

func Test_UtilSuite(t *testing.T) {
	suite.Run(t, &utilSuite{})
}

func (u *utilSuite) Test_ReadGCD() {
	gcd := GCD(10, 5)
	require.Equal(u.T(), 5, gcd)
	gcd = GCD(2000, 45)
	require.Equal(u.T(), 5, gcd)
}

func (u *utilSuite) Test_LCM() {
	lcm, err := LCM([]int{2000, 45, 400, 234})
	require.NoError(u.T(), err)
	require.Equal(u.T(), 234000, lcm)
	lcm, err = LCM([]int{123, 456, 234})
	require.NoError(u.T(), err)
	require.Equal(u.T(), 729144, lcm)
}
