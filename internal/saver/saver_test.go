package saver_test

import (
	"context"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-suggestion-api/internal/mocks"
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
	"github.com/ozoncp/ocp-suggestion-api/internal/saver"
)

var _ = Describe("Saver", func() {
	var (
		ctrl        *gomock.Controller
		ctx, ctxC   context.Context
		cancel      context.CancelFunc
		mockFlusher *mocks.MockFlusher
		s           saver.Saver
		suggestions []models.Suggestion
		err         error
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()
		mockFlusher = mocks.NewMockFlusher(ctrl)

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

	Context("When create a new saver", func() {
		BeforeEach(func() {
			s = saver.NewSaver(ctx, 4, mockFlusher, time.Second)
		})
		AfterEach(func() {
			err = s.Close(ctx)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return a saver instance", func() {
			Expect(s).ShouldNot(BeNil())
		})
	})

	Context("When the save was successful", func() {
		BeforeEach(func() {
			s = saver.NewSaver(ctx, 4, mockFlusher, 500*time.Millisecond)
		})
		AfterEach(func() {
			err = s.Close(ctx)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should save to flusher when close", func() {
			mockFlusher.EXPECT().
				Flush(ctx, gomock.Any()).
				Return(nil, nil).
				MinTimes(1)

			for _, suggestion := range suggestions {
				err = s.Save(suggestion)
				Expect(err).NotTo(HaveOccurred())
			}
		})

		It("should save to flusher when ticker", func() {
			mockFlusher.EXPECT().
				Flush(ctx, gomock.Any()).
				Return(nil, nil).
				MinTimes(1)

			for _, suggestion := range suggestions {
				err = s.Save(suggestion)
				Expect(err).NotTo(HaveOccurred())
			}
			time.Sleep(time.Second)
		})
	})

	Context("When an error occurred", func() {
		BeforeEach(func() {
			s = saver.NewSaver(ctx, 3, mockFlusher, 5*time.Second)
		})
		AfterEach(func() {
			err = s.Close(ctx)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should failed to save through flusher because of overflow capacity", func() {
			mockFlusher.EXPECT().
				Flush(ctx, gomock.Any()).
				Return(nil, nil).
				AnyTimes()
			for _, suggestion := range suggestions {
				err = s.Save(suggestion)
			}
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(saver.ErrCapacityOverflow))
		})
	})

	Context("When context cancel", func() {
		BeforeEach(func() {
			ctxC, cancel = context.WithCancel(ctx)
			s = saver.NewSaver(ctxC, 4, mockFlusher, 5*time.Second)
		})

		It("should not save to flusher", func() {
			mockFlusher.EXPECT().
				Flush(ctxC, gomock.Any()).
				Return(nil, nil).
				Times(0)

			for _, suggestion := range suggestions {
				err = s.Save(suggestion)
				Expect(err).NotTo(HaveOccurred())
			}
			cancel()
		})
		It("should panic when Close", func() {
			mockFlusher.EXPECT().
				Flush(ctxC, gomock.Any()).
				Return(nil, nil).
				AnyTimes()

			for _, suggestion := range suggestions {
				err = s.Save(suggestion)
				Expect(err).NotTo(HaveOccurred())
			}
			cancel()
			closeF := func() {
				_ = s.Close(ctxC)
			}
			time.Sleep(time.Second)
			Expect(closeF).Should(Panic())
		})
	})
})
