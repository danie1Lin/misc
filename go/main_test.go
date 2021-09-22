package main

import (
	//"github.com/stretchr/testify/assert"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ArrayVsSpliceSuite struct {
	suite.Suite
	ExpectCap int
}

func (suite *ArrayVsSpliceSuite) SetupSuite() {}
func (suite *ArrayVsSpliceSuite) SetupTest() {
	suite.ExpectCap = 4
}
func (suite *ArrayVsSpliceSuite) TearDownTest()  {}
func (suite *ArrayVsSpliceSuite) TearDownSuite() {}

func (suite *ArrayVsSpliceSuite) TestArrayInit() {
	array := [...]float64{7.0, 8.5, 9.1}
	suite.Equal(suite.ExpectCap, cap(array))
	suite.IsType(array, [3]float32{})
	suite.IsType(array, []float32{})
}

func (suite *ArrayVsSpliceSuite) TestSpliceInit() {
	array := []float64{7.0, 8.5, 9.1}
	suite.Equal(suite.ExpectCap, cap(array))
	suite.IsType(array, [3]float32{})
	suite.IsType(array, []float32{})
}

func (suite *ArrayVsSpliceSuite) AfterTest(suiteName, testName string) {
	suite.T().Log("after test", suiteName, testName)
}

func (suite *ArrayVsSpliceSuite) BeforeTest(suiteName, testName string) {
	suite.T().Log("before test", suiteName, testName)
}

func (suite *ArrayVsSpliceSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	suite.T().Log("test info", suiteName, fmt.Sprintf("%v", stats))

}

func TestArraySuite(t *testing.T) {
	suite.Run(t, new(ArrayVsSpliceSuite))
}
