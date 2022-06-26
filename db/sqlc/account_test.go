package db

import (
	"context"
	"fmt"
	util "go_finance/utils"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	category := createRandomCategory(t)
	accounts2s1 := CreateAccountsParams{
		UserID:      category.UserID,
		CategoryID:  category.ID,
		Title:       util.RandomString(12),
		Type:        category.Type,
		Description: util.RandomString(20),
		Value:       10,
		Date:        time.Now(),
	}

	accounts2, err := testQueries.CreateAccounts(context.Background(), accounts2s1)
	require.NoError(t, err)
	require.NotEmpty(t, accounts2)

	require.Equal(t, accounts2s1.UserID, accounts2.UserID)
	require.Equal(t, accounts2s1.CategoryID, accounts2.CategoryID)
	require.Equal(t, accounts2s1.Value, accounts2.Value)
	require.Equal(t, accounts2s1.Title, accounts2.Title)
	require.Equal(t, accounts2s1.Type, accounts2.Type)
	require.Equal(t, accounts2s1.Description, accounts2.Description)

	require.NotEmpty(t, accounts2.CreatedAt)
	require.NotEmpty(t, accounts2.Date)

	return accounts2
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
	fmt.Println("criada conta, categoria e usuário")
}

func TestGetAccount(t *testing.T) {
	accounts1 := createRandomAccount(t)
	accounts2, err := testQueries.GetAccount(context.Background(), accounts1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, accounts2)

	require.Equal(t, accounts1.UserID, accounts2.UserID)
	require.Equal(t, accounts1.CategoryID, accounts2.CategoryID)
	require.Equal(t, accounts1.Value, accounts2.Value)
	require.Equal(t, accounts1.Title, accounts2.Title)
	require.Equal(t, accounts1.Type, accounts2.Type)
	require.Equal(t, accounts1.Description, accounts2.Description)
}

func TestDeleteAccount(t *testing.T) {
	accounts2 := createRandomAccount(t)
	err := testQueries.DeleteCategories(context.Background(), accounts2.ID)
	require.NoError(t, err)
}

func TestUpdateAccounts(t *testing.T) {
	accounts1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:          accounts1.ID,
		Title:       util.RandomString(5),
		Description: util.RandomString(5),
		Value:       15,
	}
	accounts2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts2)

	require.Equal(t, accounts1.ID, accounts2.ID)
	require.Equal(t, arg.Title, accounts2.Title)
	require.Equal(t, arg.Description, accounts2.Description)
	require.Equal(t, arg.Value, accounts2.Value)
	require.Equal(t, accounts1.CreatedAt, accounts2.CreatedAt)
	require.NotEmpty(t, accounts1.CreatedAt)
}

func TestListAccounts(t *testing.T) {
	lastAccount := createRandomAccount(t)

	arg := GetAccountsParams{
		UserID:      lastAccount.UserID,
		CategoryID:  lastAccount.CategoryID,
		Type:        lastAccount.Type,
		Title:       lastAccount.Title,
		Description: lastAccount.Description,
		Date:        lastAccount.Date,
	}
	accounts, err := testQueries.GetAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.Equal(t, lastAccount.ID, account.ID)
		require.Equal(t, lastAccount.UserID, account.UserID)
		require.Equal(t, lastAccount.Title, account.Title)
		require.Equal(t, lastAccount.Type, account.Type)
		require.Equal(t, lastAccount.Description, account.Description)
		require.Equal(t, lastAccount.Value, account.Value)
		require.NotEmpty(t, lastAccount.CreatedAt)
		require.NotEmpty(t, lastAccount.Date)
		log.Fatal("account category title: ", account.CategoryTitle)

	}
}

func TestListGetReports(t *testing.T) {
	lastAccount := createRandomAccount(t)

	arg := GetAccountsReportsParams{
		UserID: lastAccount.UserID,
		Type:   lastAccount.Type,
	}
	sumValue, err := testQueries.GetAccountsReports(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sumValue)
}

func TestListGetGraph(t *testing.T) {
	lastAccount := createRandomAccount(t)

	arg := GetAccountsGraphParams{
		UserID: lastAccount.UserID,
		Type:   lastAccount.Type,
	}
	graphValue, err := testQueries.GetAccountsGraph(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, graphValue)
}
