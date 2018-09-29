package resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
)

type ElasticacheReplicationGroup struct {
	svc     *elasticache.ElastiCache
	groupID *string
}

func init() {
	register("ElasticacheReplicationGroup", ListElasticacheReplicationGroups)
}

func ListElasticacheReplicationGroups(sess *session.Session) ([]Resource, error) {
	svc := elasticache.New(sess)

	params := &elasticache.DescribeReplicationGroupsInput{MaxRecords: aws.Int64(100)}
	resp, err := svc.DescribeReplicationGroups(params)
	if err != nil {
		return nil, err
	}

	var resources []Resource

	for _, replicationGroup := range resp.ReplicationGroups {
		resources = append(resources, &ElasticacheReplicationGroup{
			svc:     svc,
			groupID: replicationGroup.ReplicationGroupId,
		})
	}

	return resources, nil
}

func (i *ElasticacheReplicationGroup) Remove() error {
	params := &elasticache.DeleteReplicationGroupInput{
		ReplicationGroupId: i.groupID,
	}

	_, err := i.svc.DeleteReplicationGroup(params)
	if err != nil {
		return err
	}

	return nil
}

func (i *ElasticacheReplicationGroup) String() string {
	return *i.groupID
}
