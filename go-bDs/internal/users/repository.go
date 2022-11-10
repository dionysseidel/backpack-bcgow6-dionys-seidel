package users

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/bootcamp-go/go-bDs/internal/domain"
	"github.com/google/uuid"
)

type Repository interface {
	Delete(ctx context.Context, id string) error
	GetOne(ctx context.Context, id string) (*domain.User, error)
	Store(ctx context.Context, user *domain.User) (string, error)
	Update(ctx context.Context, user domain.User) error
}

type repository struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

func NewRepository(db *dynamodb.DynamoDB) Repository {
	return &repository{
		dynamo: db,
		table:  TABLE_NAME,
	}
}

const (
	TABLE_NAME = "Users"
)

func (r *repository) Store(ctx context.Context, user *domain.User) (string, error) {
	user.Id = uuid.New().String()

	marshalledUser, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return "", err
	}

	input := &dynamodb.PutItemInput{
		Item:      marshalledUser,
		TableName: aws.String(r.table),
	}

	_, err = r.dynamo.PutItemWithContext(ctx, input)
	if err != nil {
		return "", err
	}

	return user.Id, nil
}

func (r *repository) GetOne(ctx context.Context, id string) (*domain.User, error) {
	result, err := r.dynamo.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	return domain.ItemToUser(result.Item)
}

func (r *repository) Delete(ctx context.Context, id string) error {
	_, err := r.dynamo.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(ctx context.Context, user domain.User) error {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#FN": aws.String("firstname"),
			"#LN": aws.String("lastname"),
			"#UN": aws.String("username"),
			"#PW": aws.String("password"),
			"#EM": aws.String("email"),
			"#IP": aws.String("ip"),
			"#MA": aws.String("macAddress"),
			"#WS": aws.String("website"),
			"#IM": aws.String("image"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":fn": {
				S: aws.String(user.Firstname),
			},
			":ln": {
				S: aws.String(user.Lastname),
			},
			":un": {
				S: aws.String(user.Lastname),
			},
			":pw": {
				S: aws.String(user.Password),
			},
			":em": {
				S: aws.String(user.Email),
			},
			":ip": {
				S: aws.String(user.IP),
			},
			":ma": {
				S: aws.String(user.MacAddress),
			},
			":ws": {
				S: aws.String(user.Website),
			},
			":im": {
				S: aws.String(user.Image),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(user.Id),
			},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		TableName:        aws.String(r.table),
		UpdateExpression: aws.String("SET #FN = :fn, #LN = :ln, #UN = :un, #PW = :pw, #EM = :em, #IP = :ip, #MA = :ma, #WS = :ws, #IM = :im"),
	}

	result, err := r.dynamo.UpdateItemWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				fmt.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeTransactionConflictException:
				fmt.Println(dynamodb.ErrCodeTransactionConflictException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
				return aerr
			}
		} else {
			fmt.Println(err.Error())
			return err
		}
	}

	fmt.Println(result)
	return nil
}
