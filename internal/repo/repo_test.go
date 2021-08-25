package repo_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-suggestion-api/internal/models"
	"github.com/ozoncp/ocp-suggestion-api/internal/repo"
)

var errDatabase = errors.New("database error")

var _ = Describe("Repo", func() {
	var (
		db          *sql.DB
		sqlxDB      *sqlx.DB
		sqlMock     sqlmock.Sqlmock
		err         error
		ctx         context.Context
		r           repo.Repo
		suggestions []models.Suggestion
	)

	BeforeEach(func() {
		ctx = context.Background()
		db, sqlMock, err = sqlmock.New()
		Expect(err).NotTo(HaveOccurred())
		sqlxDB = sqlx.NewDb(db, "sqlmock")
		r = repo.NewRepo(sqlxDB)

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

	Describe("Create a suggestion in repository", func() {
		Context("When create successfully", func() {
			var createdID uint64
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				sqlMock.ExpectQuery("INSERT INTO suggestions").
					WithArgs(suggestions[0].UserID, suggestions[0].CourseID).
					WillReturnRows(rows)

				createdID, err = r.CreateSuggestion(ctx, suggestions[0])
			})

			It("should should return the created ID correctly", func() {
				Expect(createdID).Should(Equal(suggestions[0].ID))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("When fails to create", func() {
			var createdID uint64
			BeforeEach(func() {
				sqlMock.ExpectQuery("INSERT INTO suggestions").
					WithArgs(suggestions[0].UserID, suggestions[0].CourseID).
					WillReturnError(errDatabase)

				createdID, err = r.CreateSuggestion(ctx, suggestions[0])
			})

			It("should return a zero ID", func() {
				Expect(createdID).Should(BeEquivalentTo(0))
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Add several suggestions to repository", func() {
		Context("When add successfully", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("INSERT INTO suggestions").
					WithArgs(
						suggestions[0].UserID, suggestions[0].CourseID,
						suggestions[1].UserID, suggestions[1].CourseID,
						suggestions[2].UserID, suggestions[2].CourseID,
						suggestions[3].UserID, suggestions[3].CourseID,
					).
					WillReturnResult(
						sqlmock.NewResult(int64(suggestions[3].UserID), int64(len(suggestions))),
					)
				err = r.AddSuggestions(ctx, suggestions)
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("When fails to add", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("INSERT INTO suggestions").
					WithArgs(
						suggestions[0].UserID, suggestions[0].CourseID,
						suggestions[1].UserID, suggestions[1].CourseID,
						suggestions[2].UserID, suggestions[2].CourseID,
						suggestions[3].UserID, suggestions[3].CourseID,
					).
					WillReturnError(errDatabase)

				err = r.AddSuggestions(ctx, suggestions)
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Describe suggestion by id", func() {
		Context("When describe successfully", func() {
			var suggestion *models.Suggestion
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "course_id"}).
					AddRow(suggestions[0].ID, suggestions[0].UserID, suggestions[0].CourseID)
				sqlMock.ExpectQuery("SELECT id, user_id, course_id FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnRows(rows)

				suggestion, err = r.DescribeSuggestion(ctx, suggestions[0].ID)
			})

			It("should return the correct suggestion", func() {
				Expect(*suggestion).Should(Equal(suggestions[0]))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("When fails to describe", func() {
			var suggestion *models.Suggestion
			BeforeEach(func() {
				sqlMock.ExpectQuery("SELECT id, user_id, course_id FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnError(errDatabase)

				suggestion, err = r.DescribeSuggestion(ctx, suggestions[0].ID)
			})

			It("should return an empty suggestion", func() {
				Expect(suggestion).Should(BeNil())
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("When not found what to describe", func() {
			var suggestion *models.Suggestion
			BeforeEach(func() {
				sqlMock.ExpectQuery("SELECT id, user_id, course_id FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnError(sql.ErrNoRows)

				suggestion, err = r.DescribeSuggestion(ctx, suggestions[0].ID)
			})

			It("should return an empty suggestion", func() {
				Expect(suggestion).Should(BeNil())
			})
			It("should error ErrSuggestionNotFound", func() {
				Expect(errors.Is(err, repo.ErrSuggestionNotFound)).Should(BeTrue())
			})
		})
	})

	Describe("List suggestions from repository", func() {
		Context("When list is successful", func() {
			const (
				maxLimit = 5
				offset   = 0
			)

			for limit := 1; limit <= maxLimit; limit++ {
				limit := limit
				var result []models.Suggestion
				BeforeEach(func() {
					rows := sqlmock.NewRows([]string{"id", "user_id", "course_id"})
					for i := 0; i < limit && i < len(suggestions); i++ {
						rows.AddRow(suggestions[i].ID, suggestions[i].UserID, suggestions[i].CourseID)
					}

					sqlMock.ExpectQuery(
						fmt.Sprintf("SELECT id, user_id, course_id FROM suggestions LIMIT %d OFFSET %d", limit, offset),
					).WillReturnRows(rows)

					result, err = r.ListSuggestions(ctx, uint64(limit), offset)
				})

				It("should return the correct result according to the limit", func() {
					Expect(result).Should(Equal(suggestions[:len(result)]))
				})
				It("should not error", func() {
					Expect(err).NotTo(HaveOccurred())
				})
			}
		})

		Context("When list fails", func() {
			const (
				limit  = 4
				offset = 0
			)
			var result []models.Suggestion
			BeforeEach(func() {
				sqlMock.ExpectQuery(
					fmt.Sprintf("SELECT id, user_id, course_id FROM suggestions LIMIT %d OFFSET %d", limit, offset),
				).WillReturnError(errDatabase)

				result, err = r.ListSuggestions(ctx, limit, offset)
			})

			It("should return the empty result", func() {
				Expect(result).Should(BeNil())
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Update a suggestion in repository", func() {
		var suggestionUpdate models.Suggestion
		BeforeEach(func() {
			suggestionUpdate = models.Suggestion{
				ID:       1,
				UserID:   44,
				CourseID: 33,
			}
		})

		Context("When update successfully", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("UPDATE suggestions SET").
					WithArgs(
						suggestionUpdate.UserID,
						suggestionUpdate.CourseID,
						suggestionUpdate.ID,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				err = r.UpdateSuggestion(ctx, suggestionUpdate)
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("When fails to update", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("UPDATE suggestions SET").
					WithArgs(
						suggestionUpdate.UserID,
						suggestionUpdate.CourseID,
						suggestionUpdate.ID,
					).
					WillReturnError(errDatabase)
				err = r.UpdateSuggestion(ctx, suggestionUpdate)
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("When not found what to update", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("UPDATE suggestions SET").
					WithArgs(
						suggestionUpdate.UserID,
						suggestionUpdate.CourseID,
						suggestionUpdate.ID,
					).
					WillReturnResult(sqlmock.NewResult(0, 0))
				err = r.UpdateSuggestion(ctx, suggestionUpdate)
			})

			It("should error ErrSuggestionNotFound", func() {
				Expect(errors.Is(err, repo.ErrSuggestionNotFound)).Should(BeTrue())
			})
		})
	})

	Describe("Remove a suggestion from repository", func() {
		Context("When remove successfully", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("DELETE FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				err = r.RemoveSuggestion(ctx, suggestions[0].ID)
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("When fails to remove", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("DELETE FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnError(errDatabase)
				err = r.RemoveSuggestion(ctx, suggestions[0].ID)
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("When not found what to remove", func() {
			BeforeEach(func() {
				sqlMock.ExpectExec("DELETE FROM suggestions WHERE").
					WithArgs(suggestions[0].ID).
					WillReturnResult(sqlmock.NewResult(0, 0))
				err = r.RemoveSuggestion(ctx, suggestions[0].ID)
			})

			It("should error ErrSuggestionNotFound", func() {
				Expect(errors.Is(err, repo.ErrSuggestionNotFound)).Should(BeTrue())
			})
		})
	})
})
