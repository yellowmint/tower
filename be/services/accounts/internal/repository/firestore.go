package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/google/uuid"
	"time"
)

type FirestoreAccountRepo struct {
	firestoreClient *firestore.Client
}

func NewFirestoreAccountRepo(firestoreClient *firestore.Client) *FirestoreAccountRepo {
	if firestoreClient == nil {
		panic("empty firestoreClient")
	}

	return &FirestoreAccountRepo{firestoreClient}
}

func (f FirestoreAccountRepo) GetAccountById(ctx context.Context, accountId uuid.UUID) (res AccountRecord, err error) {
	doc, err := f.accountById(ctx, accountId)
	if err != nil {
		return AccountRecord{}, err
	}

	err = doc.DataTo(&res)
	if err != nil {
		return AccountRecord{}, err
	}

	res.recordId = doc.Ref.ID
	return res, nil
}

func (f FirestoreAccountRepo) GetAccountByAuthUserId(ctx context.Context, authUserId string) (res AccountRecord, err error) {
	doc, err := f.accountByAuthUserId(ctx, authUserId)
	if err != nil {
		return AccountRecord{}, err
	}

	err = doc.DataTo(&res)
	if err != nil {
		return AccountRecord{}, err
	}

	res.recordId = doc.Ref.ID
	return res, nil
}

func (f FirestoreAccountRepo) CreateAccount(ctx context.Context, authUserId string, name string) error {
	_, err := f.accountByAuthUserId(ctx, authUserId)
	if err == nil {
		return ErrAccountAlreadyCreated
	}
	if err != ErrAccountNotFound {
		return err
	}

	player := AccountRecord{
		AccountId:  uuid.New().String(),
		AuthUserId: authUserId,
		Name:       name,
		CreatedAt:  time.Now().UTC(),
	}

	_, _, err = f.accountsCollection().Add(ctx, player)

	return err
}

func (f FirestoreAccountRepo) DeleteAccountById(ctx context.Context, accountId uuid.UUID) error {
	doc, err := f.accountById(ctx, accountId)
	if err == ErrAccountNotFound {
		return nil
	}
	if err != nil {
		return err
	}

	update := []firestore.Update{{
		Path:  "DeletedAt",
		Value: time.Now().UTC(),
	}}

	_, err = doc.Ref.Update(ctx, update)
	if err != nil {
		return err
	}

	return nil
}

func (f FirestoreAccountRepo) accountsCollection() *firestore.CollectionRef {
	return f.firestoreClient.Collection("accounts:accounts")
}

func (f FirestoreAccountRepo) accountByAuthUserId(ctx context.Context, authUserId string) (*firestore.DocumentSnapshot, error) {
	docs, err := f.accountsCollection().
		Where("AuthUserId", "==", authUserId).
		Where("DeletedAt", "==", nil).
		Limit(1).
		Documents(ctx).
		GetAll()

	if err != nil {
		return nil, err
	}

	if len(docs) == 0 {
		return nil, ErrAccountNotFound
	}

	return docs[0], nil
}

func (f FirestoreAccountRepo) accountById(ctx context.Context, accountId uuid.UUID) (*firestore.DocumentSnapshot, error) {
	docs, err := f.accountsCollection().
		Where("AccountId", "==", accountId.String()).
		Where("DeletedAt", "==", nil).
		Limit(1).
		Documents(ctx).
		GetAll()

	if err != nil {
		return nil, err
	}

	if len(docs) == 0 {
		return nil, ErrAccountNotFound
	}

	return docs[0], nil
}
