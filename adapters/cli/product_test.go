package cli_test

import (
 "testing"
 "github.com/golang/mock/gomock"
 "github.com/satori/go.uuid"
 "fmt"
 mock_application "github.com/alessandroprudencio/Go-Hexagonal/application/mocks"
 "github.com/alessandroprudencio/Go-Hexagonal/adapters/cli"
 "github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {

    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    productName := "Product Test"
    productPrice := 23.10
    productStatus := "enabled"
    productId := uuid.NewV4().String()

    productMock := mock_application.NewMockProductInterface(ctrl)

    productMock.EXPECT().GetID().Return(productId).AnyTimes()
    productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
    productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
    productMock.EXPECT().GetName().Return(productName).AnyTimes()

    service := mock_application.NewMockProductServiceInterface(ctrl)

    service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
    service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
    service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
    service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

    resultExpected := fmt.Sprintf(
        "Product id  %s with the name %s has been created with price %f and status %s",
        productId,
        productName,
        productPrice,
        productStatus,
    )

    result, err := cli.Run(service, "create", "", productName, productPrice)
    require.Nil(t, err)
    require.Equal(t, resultExpected, result)

     resultExpected = fmt.Sprintf("Product  %s has been enabled.",
        productName)

    result, err = cli.Run(service, "enable", productId, "", 0)
    require.Nil(t, err)
    require.Equal(t, resultExpected, result)

    resultExpected = fmt.Sprintf("Product  %s has been disabled.",
        productName)

    result, err = cli.Run(service, "disable", productId, "", 0)
    require.Nil(t, err)
    require.Equal(t, resultExpected, result)

    resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
        productId, productName, productPrice, productStatus)

    result, err = cli.Run(service, "get", productId, "", 0)
    require.Nil(t, err)
    require.Equal(t, resultExpected, result)
}
