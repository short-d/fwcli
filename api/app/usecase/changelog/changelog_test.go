// +build !integration all

package changelog

import (
	"testing"

	"github.com/short-d/app-template/backend/app/entity"
	"github.com/short-d/app-template/backend/app/usecase/repo"
	"github.com/short-d/app/fw/assert"
)

func TestChangeLog_GetChangeLog(t *testing.T) {
	t.Parallel()

	summaryMarkdown1 := "summary 1"
	summaryMarkdown2 := "summary 2"
	testCases := []struct {
		name      string
		changeLog []entity.Change
	}{
		{
			name: "get full changelog successfully",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
		}, {
			name:      "get empty changelog successfully",
			changeLog: []entity.Change{},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repo.NewChangeLogFake(testCase.changeLog)
			cl := NewChangeLog(&changeLogRepo)

			changeLog, err := cl.GetChangeLog()
			assert.Equal(t, nil, err)
			assert.SameElements(t, testCase.changeLog, changeLog)
		})
	}
}
