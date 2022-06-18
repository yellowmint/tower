package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func (f *FirestoreAccountRepo) GetAccountById(ctx context.Context, accountId uuid.UUID) (res AccountRecord, err error) {
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

func (f *FirestoreAccountRepo) GetAccountByAuthUserId(ctx context.Context, authUserId string) (res AccountRecord, err error) {
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

func (f *FirestoreAccountRepo) GetNameCounter(ctx context.Context, name string) (res NameCounterRecord, err error) {
	doc, err := f.nameCounterByName(ctx, name)
	if err != nil {
		return NameCounterRecord{}, err
	}

	err = doc.DataTo(&res)
	if err != nil {
		return NameCounterRecord{}, err
	}

	return res, nil
}

func (f *FirestoreAccountRepo) UpdateNameCounter(ctx context.Context, nameCounter NameCounterRecord) (err error) {
	doc, err := f.nameCounterByName(ctx, nameCounter.Name)
	if err == ErrNameCounterNotFound {
		_, _, err := f.namesCounterCollection().Add(ctx, nameCounter)
		if err != nil {
			return errors.Wrap(err, "insert name counter")
		}

		return nil
	}
	if err != nil {
		return err
	}

	updates := []firestore.Update{{
		Path:  "Count",
		Value: nameCounter.Count,
	}}

	_, err = doc.Ref.Update(ctx, updates)
	if err != nil {
		return errors.Wrap(err, "update name counter")
	}

	return nil
}

func (f *FirestoreAccountRepo) CreateAccount(ctx context.Context, authUserId string, record AccountRecord) error {
	err := f.UpdateNameCounter(ctx, NameCounterRecord{Name: record.Name, Count: record.NameNumber})
	if err != nil {
		return err
	}

	_, err = f.accountByAuthUserId(ctx, authUserId)
	if err == nil {
		return ErrAccountAlreadyCreated
	}
	if err != ErrAccountNotFound {
		return err
	}

	record.AuthUserId = authUserId
	record.CreatedAt = time.Now().UTC()

	_, _, err = f.accountsCollection().Add(ctx, record)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreAccountRepo) DeleteAccountById(ctx context.Context, accountId uuid.UUID) error {
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

func (f *FirestoreAccountRepo) accountsCollection() *firestore.CollectionRef {
	return f.firestoreClient.Collection("accounts:accounts")
}

func (f *FirestoreAccountRepo) accountByAuthUserId(ctx context.Context, authUserId string) (*firestore.DocumentSnapshot, error) {
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

func (f *FirestoreAccountRepo) accountById(ctx context.Context, accountId uuid.UUID) (*firestore.DocumentSnapshot, error) {
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

func (f *FirestoreAccountRepo) namesCounterCollection() *firestore.CollectionRef {
	return f.firestoreClient.Collection("accounts:names_counter")
}

func (f *FirestoreAccountRepo) nameCounterByName(ctx context.Context, name string) (*firestore.DocumentSnapshot, error) {
	docs, err := f.namesCounterCollection().
		Where("Name", "==", name).
		Limit(1).
		Documents(ctx).
		GetAll()

	if err != nil {
		return nil, err
	}

	if len(docs) == 0 {
		return nil, ErrNameCounterNotFound
	}

	return docs[0], nil
}
