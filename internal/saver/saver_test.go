package saver_test

import (
	"context"
	"errors"
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
		ctx         context.Context
		mockFlusher *mocks.MockFlusher
		suggestions []models.Suggestion
		err error
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
		It("should return a saver instance", func() {
			s := saver.NewSaver(4, mockFlusher, 5*time.Second)
			Expect(s).ShouldNot(BeNil())
		})

		It("can make multiple Init() and Close()", func() {
			s := saver.NewSaver(4, mockFlusher, 5*time.Second)
			Expect(s.IsInit()).Should(BeFalse())
			Expect(s.IsClosed()).Should(BeTrue())

			s.Init(ctx)
			Expect(s.IsInit()).Should(BeTrue())
			Expect(s.IsClosed()).Should(BeFalse())
			err = s.Close(ctx)
			Expect(s.IsClosed()).Should(BeTrue())
			Expect(err).NotTo(HaveOccurred())

			s.Init(ctx)
			Expect(s.IsInit()).Should(BeTrue())
			s.Init(ctx)
			Expect(s.IsInit()).Should(BeTrue())
			err = s.Close(ctx)
			Expect(s.IsClosed()).Should(BeTrue())
			Expect(err).NotTo(HaveOccurred())
			err = s.Close(ctx)
			Expect(s.IsClosed()).Should(BeTrue())
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("When the save was successful", func() {
		It("should save to flusher when close", func() {
			mockFlusher.EXPECT().
				Flush(gomock.Any(), gomock.Any()).
				Return(nil, nil).
				MinTimes(1)

			s := saver.NewSaver(4, mockFlusher, 5*time.Second)
			//s.Init() - не делаем, чтобы не запускать таймер
			for _, suggestion := range suggestions {
				err = s.Save(ctx, suggestion)
				Expect(err).NotTo(HaveOccurred())
			}
			Expect(s.IsBufferEmpty()).Should(BeFalse()) //Внутренний буфер после заполнения - непустой

			err = s.Close(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(s.IsBufferEmpty()).Should(BeTrue()) //А после Close буфер - пустой
		})

		It("should save to flusher when ticker", func() {
			mockFlusher.EXPECT().
				Flush(gomock.Any(), gomock.Any()).
				Return(nil, nil).
				MinTimes(1)

			s := saver.NewSaver(4, mockFlusher, 100*time.Millisecond)
			s.Init(ctx)
			for _, suggestion := range suggestions {
				err = s.Save(ctx, suggestion)
				Expect(err).NotTo(HaveOccurred())
			}

			Expect(s.IsBufferEmpty()).Should(BeFalse()) //Внутренний буфер после заполнения - непустой
			time.Sleep(500 * time.Millisecond)
			Expect(s.IsBufferEmpty()).Should(BeTrue()) //А через полсекунды буфер - пустой
			err = s.Close(ctx)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should extra flush when the capacity is exceeded", func() {
			mockFlusher.EXPECT().
				Flush(gomock.Any(), gomock.Any()).
				Return(nil, nil).
				MinTimes(2) //Минимум 2 Flush: extra + Close

			const capacity = 2
			s := saver.NewSaver(capacity, mockFlusher, 5*time.Second)
			s.Init(ctx)
			for i, suggestion := range suggestions {
				err = s.Save(ctx, suggestion)
				Expect(err).NotTo(HaveOccurred())
				if i == capacity { //превышение capacity -> extra flush, после чего буфер пустой
					Expect(s.IsBufferEmpty()).Should(BeTrue())
				} else { //иначе буфер - непустой
					Expect(s.IsBufferEmpty()).Should(BeFalse())
				}
			}

			err = s.Close(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(s.IsBufferEmpty()).Should(BeTrue())
		})
	})

	Context("When an error occurred", func() {
		It("should be error when close and data left in inner buffer", func() {
			mockFlusher.EXPECT().
				Flush(gomock.Any(), suggestions).
				Return(suggestions, errors.New("error")).
				MinTimes(1)

			s := saver.NewSaver(4, mockFlusher, 5*time.Second)
			s.Init(ctx)
			for _, suggestion := range suggestions {
				err = s.Save(ctx, suggestion)
				Expect(err).NotTo(HaveOccurred())
			}
			Expect(s.IsBufferEmpty()).Should(BeFalse())
			err = s.Close(ctx)
			Expect(s.IsBufferEmpty()).Should(BeFalse())
			Expect(err).To(HaveOccurred())
		})

		It("should be error when save, data left in inner buffer and flush when close", func() {
			failCall := mockFlusher.EXPECT().
				Flush(gomock.Any(), gomock.Any()).
				Return(nil, errors.New("error"))
			successCall := mockFlusher.EXPECT().
				Flush(gomock.Any(), gomock.Any()).
				Return(nil, nil)
			gomock.InOrder(failCall, successCall)

			const capacity = 2
			s := saver.NewSaver(capacity, mockFlusher, 5*time.Second)
			s.Init(ctx)
			for i, suggestion := range suggestions {
				err = s.Save(ctx, suggestion)
				if i == capacity { //превышение capacity -> extra flush, когда должна случиться ошибка
					Expect(err).To(HaveOccurred())
				} else {
					Expect(err).NotTo(HaveOccurred())
				}
			}
			Expect(s.IsBufferEmpty()).Should(BeFalse())
			err = s.Close(ctx)
			Expect(s.IsBufferEmpty()).Should(BeTrue())
			Expect(err).NotTo(HaveOccurred())
		})

		It("should be error when ticker", func() {
			defer GinkgoRecover()
			mockFlusher.EXPECT().
				Flush(gomock.Any(), suggestions).
				Return(suggestions, errors.New("error")).
				AnyTimes()

			s := saver.NewSaver(4, mockFlusher, 100*time.Millisecond)
			s.Init(ctx)
			for _, suggestion := range suggestions {
				err = s.Save(ctx, suggestion)
				Expect(err).NotTo(HaveOccurred())
			}

			Expect(s.IsBufferEmpty()).Should(BeFalse()) //Внутренний буфер после заполнения - непустой
			time.Sleep(500 * time.Millisecond)
			Expect(s.IsBufferEmpty()).Should(BeFalse()) //И так как ошибка, остаётся непустым
			err = s.Close(ctx)
			Expect(err).To(HaveOccurred())
		})

	})

	Context("When context cancel", func() {
		It("should not save to flusher", func() {
			ctxC, cancel := context.WithCancel(ctx)
			mockFlusher.EXPECT().
				Flush(ctxC, gomock.Any()).
				Return(nil, nil).
				Times(0)

			s := saver.NewSaver(4, mockFlusher, 5*time.Second)
			s.Init(ctxC)
			for _, suggestion := range suggestions {
				err = s.Save(ctxC, suggestion)
				Expect(err).NotTo(HaveOccurred())
			}
			Expect(s.IsBufferEmpty()).Should(BeFalse())

			cancel()
			time.Sleep(500 * time.Millisecond)
			Expect(s.IsBufferEmpty()).Should(BeFalse())
		})
	})
})
