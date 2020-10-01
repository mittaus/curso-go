package integration

import (
	"testing"
)

func TestArticle_Create(t *testing.T) {

	// db, err := NewRepositories("sqlserver", "sa", "Pa$$w0rd", "1433", "localhost", "DDD_DEV")
	// if err != nil {
	// 	fmt.Println("Hola")
	// }

	// rw := ArticleRepository.New(db)

	// happyInsert := func(t *testing.T, toInsert domain.User) {
	// 	returnedUser, err := rw.Create(toInsert.Name, toInsert.Email, toInsert.Password)
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, toInsert.Name, returnedUser.Name)
	// 	assert.Equal(t, toInsert.Email, returnedUser.Email)
	// 	assert.Equal(t, toInsert.Password, returnedUser.Password)
	// }

	// faillingInsert := func(t *testing.T, toInsert domain.User) {
	// 	returnedUser, err := rw.Create(toInsert.Name, toInsert.Email, toInsert.Password)
	// 	assert.Error(t, err)
	// 	assert.Nil(t, returnedUser)
	// }

	// var wg sync.WaitGroup
	// wg.Add(2)
	// go func(t *testing.T, toInsert domain.User) {
	// 	defer wg.Done()
	// 	happyInsert(t, toInsert)
	// 	faillingInsert(t, toInsert)
	// }(t, testData.User("jane"))
	// go func(t *testing.T, toInsert domain.User) {
	// 	defer wg.Done()
	// 	happyInsert(t, toInsert)
	// 	faillingInsert(t, toInsert)
	// }(t, testData.User("rick"))
	// wg.Wait()
}
