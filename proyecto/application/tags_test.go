package application_test

import (
	"testing"

	"example.com/mittaus/ddd-example/infraestructure/application.mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInteractor_Tags(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	tags := []string{"tag1", "tag2"}

	i := mock.NewMockedInteractor(mockCtrl)
	i.TagsRW.EXPECT().GetAll().Return(tags, nil).Times(1)

	retTags, err := i.GetUCHandler().Tags()

	assert.NoError(t, err)
	assert.Equal(t, tags, retTags)
}
