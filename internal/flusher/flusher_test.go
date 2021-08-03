package flusher_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-suggestion-api/internal/flusher"
	"github.com/ozoncp/ocp-suggestion-api/internal/mocks"
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl        *gomock.Controller
		ctx         context.Context
		mockRepo    *mocks.MockRepo
		f           flusher.Flusher
		suggestions []models.Suggestion
		result      []models.Suggestion
		err         error
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()
		mockRepo = mocks.NewMockRepo(ctrl)

		suggestions = []models.Suggestion{
			{ID: 1, UserID: 1, CourseID: 1},
			{ID: 2, UserID: 2, CourseID: 2},
			{ID: 3, UserID: 3, CourseID: 3},
			{ID: 4, UserID: 4, CourseID: 4},
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("When the load was successful", func() {
		It("should upload to repo in 1 successful call, no remains left", func() {
			f = flusher.NewFlusher(4, mockRepo)
			mockRepo.EXPECT().
				AddSuggestions(ctx, gomock.Any()).
				Return(nil).
				Times(1)
			result, err = f.Flush(ctx, suggestions)
			Expect(result).Should(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should upload to repo in 2 successful calls, no remains left", func() {
			f = flusher.NewFlusher(2, mockRepo)
			mockRepo.EXPECT().
				AddSuggestions(ctx, gomock.Any()).
				Return(nil).
				Times(2)
			result, err = f.Flush(ctx, suggestions)
			Expect(result).Should(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("When an error occurred", func() {
		BeforeEach(func() {
			f = flusher.NewFlusher(3, mockRepo)
		})

		It("should failed to upload to repo, all items remain", func() {
			mockRepo.EXPECT().
				AddSuggestions(ctx, gomock.Any()).
				Return(errors.New("error")).
				Times(1)

			result, err = f.Flush(ctx, suggestions)
			Expect(result).To(Equal(suggestions))
			Expect(err).To(HaveOccurred())
		})

		It("should upload to repo 1st chunk (1 successful call), then 1 failure call and 1 item remains", func() {
			successCall := mockRepo.EXPECT().
				AddSuggestions(ctx, gomock.Any()).
				Return(nil)
			failCall := mockRepo.EXPECT().
				AddSuggestions(ctx, gomock.Any()).
				Return(errors.New("error"))
			gomock.InOrder(successCall, failCall)

			result, err = f.Flush(ctx, suggestions)
			Expect(result).To(Equal(suggestions[3:]),
				"rest that failed to upload to repo")
			Expect(err).To(HaveOccurred())
		})
	})
})
