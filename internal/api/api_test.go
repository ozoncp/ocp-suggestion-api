package api_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-suggestion-api/internal/api"
	"github.com/ozoncp/ocp-suggestion-api/internal/models"
	"github.com/ozoncp/ocp-suggestion-api/internal/repo"
	desc "github.com/ozoncp/ocp-suggestion-api/pkg/ocp-suggestion-api"
)

var errDatabase = errors.New("database error")

var _ = Describe("Api", func() {
	var (
		db          *sql.DB
		sqlxDB      *sqlx.DB
		sqlMock     sqlmock.Sqlmock
		err         error
		ctx         context.Context
		r           repo.Repo
		suggestions []models.Suggestion
		grpc        desc.OcpSuggestionApiServer
	)

	BeforeEach(func() {
		ctx = context.Background()
		db, sqlMock, err = sqlmock.New()
		Expect(err).NotTo(HaveOccurred())
		sqlxDB = sqlx.NewDb(db, "sqlmock")
		r = repo.NewRepo(sqlxDB)
		grpc = api.NewSuggestionAPI(r)

		suggestions = []models.Suggestion{
			{ID: 1, UserID: 1, CourseID: 1},
			{ID: 2, UserID: 2, CourseID: 2},
			{ID: 3, UserID: 3, CourseID: 3},
			{ID: 4, UserID: 4, CourseID: 4},
		}
	})

	AfterEach(func() {
		sqlMock.ExpectClose()
		err = db.Close()
		Expect(err).NotTo(HaveOccurred())
		Expect(sqlMock.ExpectationsWereMet()).NotTo(HaveOccurred())
	})

	Describe("Create a suggestion", func() {
		var (
			request  *desc.CreateSuggestionV1Request
			response *desc.CreateSuggestionV1Response
		)
		BeforeEach(func() {
			request = &desc.CreateSuggestionV1Request{
				UserId:   suggestions[0].UserID,
				CourseId: suggestions[0].CourseID,
			}
		})

		Context("When create successfully", func() {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				sqlMock.ExpectQuery("INSERT INTO suggestions").
					WithArgs(suggestions[0].UserID, suggestions[0].CourseID).
					WillReturnRows(rows)

				response, err = grpc.CreateSuggestionV1(ctx, request)
			})

			It("should should return the created ID correctly", func() {
				Expect(response.SuggestionId).Should(Equal(suggestions[0].ID))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("When fails to create due to invalid arguments", func() {
			BeforeEach(func() {
				request.UserId = 0

				response, err = grpc.CreateSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return an invalid argument error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
			})
		})

		Context("When fails to create due to database error", func() {
			BeforeEach(func() {
				sqlMock.ExpectQuery("INSERT INTO suggestions").
					WithArgs(suggestions[0].UserID, suggestions[0].CourseID).
					WillReturnError(errDatabase)

				response, err = grpc.CreateSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return an internal error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.Internal))
			})
		})
	})

	Describe("Describe suggestion", func() {
		var (
			request  *desc.DescribeSuggestionV1Request
			response *desc.DescribeSuggestionV1Response
		)
		BeforeEach(func() {
			request = &desc.DescribeSuggestionV1Request{
				SuggestionId: suggestions[0].ID,
			}
		})

		Context("When describe successfully", func() {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "course_id"}).
					AddRow(suggestions[0].ID, suggestions[0].UserID, suggestions[0].CourseID)
				sqlMock.ExpectQuery("SELECT id, user_id, course_id FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnRows(rows)

				response, err = grpc.DescribeSuggestionV1(ctx, request)
			})

			It("should return the correct suggestion", func() {
				Expect(response.Suggestion.UserId).Should(Equal(suggestions[0].UserID))
				Expect(response.Suggestion.CourseId).Should(Equal(suggestions[0].CourseID))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("When fails to describe due to database error", func() {
			BeforeEach(func() {
				sqlMock.ExpectQuery("SELECT id, user_id, course_id FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnError(errDatabase)

				response, err = grpc.DescribeSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return an internal error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.Internal))
			})
		})

		Context("When not found what to describe", func() {
			BeforeEach(func() {
				sqlMock.ExpectQuery("SELECT id, user_id, course_id FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnError(sql.ErrNoRows)

				response, err = grpc.DescribeSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return a not found error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.NotFound))
			})
		})

		Context("When fails to create due to invalid arguments", func() {
			BeforeEach(func() {
				request.SuggestionId = 0

				response, err = grpc.DescribeSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return an invalid argument error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
			})
		})
	})

	Describe("List suggestions", func() {
		const (
			maxLimit = 5
			offset   = 0
		)
		var (
			request  *desc.ListSuggestionV1Request
			response *desc.ListSuggestionV1Response
		)

		Context("When list is successful", func() {
			for limit := 1; limit <= maxLimit; limit++ {
				limit := limit
				BeforeEach(func() {
					request = &desc.ListSuggestionV1Request{
						Limit:  uint64(limit),
						Offset: uint64(offset),
					}
					rows := sqlmock.NewRows([]string{"id", "user_id", "course_id"})
					for i := 0; i < limit && i < len(suggestions); i++ {
						rows.AddRow(suggestions[i].ID, suggestions[i].UserID, suggestions[i].CourseID)
					}
					sqlMock.ExpectQuery(
						fmt.Sprintf("SELECT id, user_id, course_id FROM suggestions LIMIT %d OFFSET %d", limit, offset),
					).WillReturnRows(rows)

					response, err = grpc.ListSuggestionV1(ctx, request)
				})

				It("should return the correct result according to the limit", func() {
					for i := 0; i < len(response.Suggestions); i++ {
						Expect(response.Suggestions[i].Id).Should(Equal(suggestions[i].ID))
						Expect(response.Suggestions[i].UserId).Should(Equal(suggestions[i].UserID))
						Expect(response.Suggestions[i].CourseId).Should(Equal(suggestions[i].CourseID))
					}
				})
				It("should not error", func() {
					Expect(err).NotTo(HaveOccurred())
				})
			}
		})

		Context("When list fails due to database error", func() {
			BeforeEach(func() {
				const (
					limit  = 4
					offset = 0
				)
				sqlMock.ExpectQuery(
					fmt.Sprintf("SELECT id, user_id, course_id FROM suggestions LIMIT %d OFFSET %d", limit, offset),
				).WillReturnError(errDatabase)

				request = &desc.ListSuggestionV1Request{
					Limit:  uint64(limit),
					Offset: uint64(offset),
				}

				response, err = grpc.ListSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return an internal error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.Internal))
			})
		})

		Context("When list fails due to invalid arguments", func() {
			BeforeEach(func() {
				request = &desc.ListSuggestionV1Request{
					Limit:  0,
					Offset: 0,
				}

				response, err = grpc.ListSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return an invalid argument error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
			})
		})
	})

	Describe("Update a suggestion", func() {
		var (
			request  *desc.UpdateSuggestionV1Request
			response *desc.UpdateSuggestionV1Response
		)
		BeforeEach(func() {
			request = &desc.UpdateSuggestionV1Request{
				Suggestion: &desc.Suggestion{
					Id:       1,
					UserId:   44,
					CourseId: 33,
				},
			}
		})

		Context("When update successfully", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("UPDATE suggestions SET").
					WithArgs(
						request.Suggestion.UserId,
						request.Suggestion.CourseId,
						request.Suggestion.Id,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))

				response, err = grpc.UpdateSuggestionV1(ctx, request)
			})

			It("should be an empty response", func() {
				Expect(response).Should(Equal(&desc.UpdateSuggestionV1Response{}))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("When fails to update due to database error", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("UPDATE suggestions SET").
					WithArgs(
						request.Suggestion.UserId,
						request.Suggestion.CourseId,
						request.Suggestion.Id,
					).
					WillReturnError(errDatabase)

				response, err = grpc.UpdateSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return an internal error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.Internal))
			})
		})

		Context("When not found what to update", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("UPDATE suggestions SET").
					WithArgs(
						request.Suggestion.UserId,
						request.Suggestion.CourseId,
						request.Suggestion.Id,
					).
					WillReturnResult(sqlmock.NewResult(0, 0))

				response, err = grpc.UpdateSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return a not found error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.NotFound))
			})
		})

		Context("When fails to update due to invalid arguments", func() {
			BeforeEach(func() {
				request.Suggestion.Id = 0

				response, err = grpc.UpdateSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return an invalid argument error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
			})
		})
	})

	Describe("Remove a suggestion", func() {
		var (
			request  *desc.RemoveSuggestionV1Request
			response *desc.RemoveSuggestionV1Response
		)
		BeforeEach(func() {
			request = &desc.RemoveSuggestionV1Request{
				SuggestionId: 1,
			}
		})

		Context("When remove successfully", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("DELETE FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnResult(sqlmock.NewResult(1, 1))

				response, err = grpc.RemoveSuggestionV1(ctx, request)
			})

			It("should be an empty response", func() {
				Expect(response).Should(Equal(&desc.RemoveSuggestionV1Response{}))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("When fails to remove due to database error", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("DELETE FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnError(errDatabase)

				response, err = grpc.RemoveSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return an internal error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.Internal))
			})
		})

		Context("When not found what to remove", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("DELETE FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnResult(sqlmock.NewResult(0, 0))

				response, err = grpc.RemoveSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return a not found error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.NotFound))
			})
		})

		Context("When fails to remove due to invalid arguments", func() {
			BeforeEach(func() {
				request.SuggestionId = 0

				response, err = grpc.RemoveSuggestionV1(ctx, request)
			})

			It("should be nil response", func() {
				Expect(response).Should(BeNil())
			})
			It("should return an invalid argument error", func() {
				Expect(status.Convert(err).Code()).Should(Equal(codes.InvalidArgument))
			})
		})
	})
})
