package unitofwork_test

import (
	"context"
	"errors"
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gomock "go.uber.org/mock/gomock"

	database_mocks "github.com/sdv-projects/go-ea-common/database/mocks"
	unitofwork "github.com/sdv-projects/go-ea-common/unitofwork"
	unitofwork_mocks "github.com/sdv-projects/go-ea-common/unitofwork/mocks"
)

type utEntity1 struct {
	ID   string
	Name string
}

type utEntity2 struct {
	ID   string
	Type string
}

var _ = Describe("Unit of work (uow) implemetation tests", func() {
	var (
		txFactoryMock    *database_mocks.MockTxFactory
		txMock           *database_mocks.MockTx
		reposFactoryMock *unitofwork_mocks.MockUoWRepositoryFactory
		reposMock        *unitofwork_mocks.MockUoWRepository

		entityNew1 = &utEntity1{ID: "1", Name: "1"}
		entityNew2 = &utEntity2{ID: "2", Type: "2"}
		entityUpd1 = &utEntity1{ID: "3", Name: "3"}
		entityUpd2 = &utEntity2{ID: "4", Type: "4"}
		entityDel1 = &utEntity1{ID: "5", Name: "5"}
		entityDel2 = &utEntity2{ID: "6", Type: "6"}
	)

	var _ = BeforeEach(func() {
		ctx = context.Background()
		txFactoryMock = database_mocks.NewMockTxFactory(ctl)
		txMock = database_mocks.NewMockTx(ctl)
		reposFactoryMock = unitofwork_mocks.NewMockUoWRepositoryFactory(ctl)
		reposMock = unitofwork_mocks.NewMockUoWRepository(ctl)
	})
	Context("Sunny", func() {
		It("Register entites and commit", func() {
			reposMock.EXPECT().TxInsert(ctx, txMock, entityNew1).Return(nil).Times(1)
			reposMock.EXPECT().TxInsert(ctx, txMock, entityNew2).Return(nil).Times(1)

			reposMock.EXPECT().TxUpdate(ctx, txMock, entityUpd1).Return(nil).Times(1)
			reposMock.EXPECT().TxUpdate(ctx, txMock, entityUpd2).Return(nil).Times(1)

			reposMock.EXPECT().TxDelete(ctx, txMock, entityDel1).Return(nil).Times(1)
			reposMock.EXPECT().TxDelete(ctx, txMock, entityDel2).Return(nil).Times(1)

			reposFactoryMock.EXPECT().
				Get(ctx, reflect.TypeOf((*utEntity1)(nil))).
				Return(reposMock, nil).
				MinTimes(1)
			reposFactoryMock.EXPECT().
				Get(ctx, reflect.TypeOf((*utEntity2)(nil))).
				Return(reposMock, nil).
				MinTimes(1)

			txMock.EXPECT().Commit().Return(nil).Times(1)
			txFactoryMock.EXPECT().BeginTx(ctx).Return(txMock, nil).Times(1)

			uow := unitofwork.NewUoW(txFactoryMock, reposFactoryMock)
			Expect(uow).ShouldNot(BeNil())

			uow.RegisterNew(entityNew1)
			uow.RegisterNew(entityNew2)
			uow.RegisterDirty(entityUpd1)
			uow.RegisterDirty(entityUpd2)
			uow.RegisterDeleted(entityDel1)
			uow.RegisterDeleted(entityDel2)

			err := uow.Commit(ctx)
			Expect(err).Should(Succeed())
		})
	})
	Context("Rainy", func() {
		It("Register nil entity", func() {
			reposMock.EXPECT().TxInsert(ctx, txMock, gomock.Any()).Return(nil).Times(0)
			reposMock.EXPECT().TxUpdate(ctx, txMock, gomock.Any()).Return(nil).Times(0)
			reposMock.EXPECT().TxDelete(ctx, txMock, gomock.Any()).Return(nil).Times(0)

			reposFactoryMock.EXPECT().
				Get(ctx, gomock.Any()).
				Return(reposMock, nil).
				Times(0)

			txMock.EXPECT().Commit().Return(nil).Times(1)
			txFactoryMock.EXPECT().BeginTx(ctx).Return(txMock, nil).Times(1)

			uow := unitofwork.NewUoW(txFactoryMock, reposFactoryMock)
			Expect(uow).ShouldNot(BeNil())

			uow.RegisterNew(nil)
			uow.RegisterDirty(nil)
			uow.RegisterDeleted(nil)

			err := uow.Commit(ctx)
			Expect(err).Should(Succeed())
		})
		It("Failed to execute TxFactory.BeginTx", func() {
			reposFactoryMock.EXPECT().
				Get(ctx, gomock.Any()).
				Return(reposMock, nil).
				AnyTimes()

			txFactoryMock.EXPECT().BeginTx(ctx).Return(nil, errors.ErrUnsupported).Times(1)

			uow := unitofwork.NewUoW(txFactoryMock, reposFactoryMock)
			Expect(uow).ShouldNot(BeNil())

			uow.RegisterNew(entityNew1)

			err := uow.Commit(ctx)
			Expect(err).ShouldNot(Succeed())
			Expect(err).To(MatchError(errors.ErrUnsupported))
		})
		It("Failed to get UoWReposirory by entity type", func() {
			reposFactoryMock.EXPECT().
				Get(ctx, gomock.Any()).
				Return(nil, errors.ErrUnsupported).
				AnyTimes()

			txMock.EXPECT().Commit().Return(nil).Times(0)
			txMock.EXPECT().Rollback().Return(nil).Times(1)
			txFactoryMock.EXPECT().BeginTx(ctx).Return(txMock, nil).Times(1)

			uow := unitofwork.NewUoW(txFactoryMock, reposFactoryMock)
			Expect(uow).ShouldNot(BeNil())

			uow.RegisterNew(entityNew1)

			err := uow.Commit(ctx)
			Expect(err).ShouldNot(Succeed())
			Expect(err).To(MatchError(errors.ErrUnsupported))
		})
		It("Failed to execute UoWRepository.TxInsert", func() {
			reposMock.EXPECT().TxInsert(ctx, txMock, gomock.Any()).Return(errors.ErrUnsupported).Times(1)

			reposFactoryMock.EXPECT().
				Get(ctx, gomock.Any()).
				Return(reposMock, nil).
				AnyTimes()

			txMock.EXPECT().Commit().Return(nil).Times(0)
			txMock.EXPECT().Rollback().Return(nil).Times(1)
			txFactoryMock.EXPECT().BeginTx(ctx).Return(txMock, nil).Times(1)

			uow := unitofwork.NewUoW(txFactoryMock, reposFactoryMock)
			Expect(uow).ShouldNot(BeNil())

			uow.RegisterNew(entityNew1)
			uow.RegisterNew(entityNew2)

			err := uow.Commit(ctx)
			Expect(err).ShouldNot(Succeed())
			Expect(err).To(MatchError(errors.ErrUnsupported))
		})
		It("Failed to execute UoWRepository.TxUpdate", func() {
			reposMock.EXPECT().TxUpdate(ctx, txMock, gomock.Any()).Return(errors.ErrUnsupported).Times(1)

			reposFactoryMock.EXPECT().
				Get(ctx, gomock.Any()).
				Return(reposMock, nil).
				AnyTimes()

			txMock.EXPECT().Commit().Return(nil).Times(0)
			txMock.EXPECT().Rollback().Return(nil).Times(1)
			txFactoryMock.EXPECT().BeginTx(ctx).Return(txMock, nil).Times(1)

			uow := unitofwork.NewUoW(txFactoryMock, reposFactoryMock)
			Expect(uow).ShouldNot(BeNil())

			uow.RegisterDirty(entityUpd1)
			uow.RegisterDirty(entityUpd2)

			err := uow.Commit(ctx)
			Expect(err).ShouldNot(Succeed())
			Expect(err).To(MatchError(errors.ErrUnsupported))
		})
		It("Failed to execute UoWRepository.TxDelete", func() {
			reposMock.EXPECT().TxDelete(ctx, txMock, gomock.Any()).Return(errors.ErrUnsupported).Times(1)

			reposFactoryMock.EXPECT().
				Get(ctx, gomock.Any()).
				Return(reposMock, nil).
				AnyTimes()

			txMock.EXPECT().Commit().Return(nil).Times(0)
			txMock.EXPECT().Rollback().Return(nil).Times(1)
			txFactoryMock.EXPECT().BeginTx(ctx).Return(txMock, nil).Times(1)

			uow := unitofwork.NewUoW(txFactoryMock, reposFactoryMock)
			Expect(uow).ShouldNot(BeNil())

			uow.RegisterDeleted(entityDel1)
			uow.RegisterDeleted(entityDel2)

			err := uow.Commit(ctx)
			Expect(err).ShouldNot(Succeed())
			Expect(err).To(MatchError(errors.ErrUnsupported))
		})
		It("Failed to execute Tx.Commit", func() {
			reposMock.EXPECT().TxInsert(ctx, txMock, gomock.Any()).Return(nil).Times(1)

			reposFactoryMock.EXPECT().
				Get(ctx, gomock.Any()).
				Return(reposMock, nil).
				AnyTimes()

			txMock.EXPECT().Commit().Return(errors.ErrUnsupported).Times(1)
			txMock.EXPECT().Rollback().Return(nil).Times(1)
			txFactoryMock.EXPECT().BeginTx(ctx).Return(txMock, nil).Times(1)

			uow := unitofwork.NewUoW(txFactoryMock, reposFactoryMock)
			Expect(uow).ShouldNot(BeNil())

			uow.RegisterNew(entityNew1)

			err := uow.Commit(ctx)
			Expect(err).ShouldNot(Succeed())
			Expect(err).To(MatchError(errors.ErrUnsupported))
		})
	})
})
