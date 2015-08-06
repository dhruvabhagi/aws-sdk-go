package elasticacheiface

import (
	"github.com/datacratic/aws-sdk-go/service/elasticache"
)

type ElastiCacheAPI interface {
	AddTagsToResource(*elasticache.AddTagsToResourceInput) (*elasticache.TagListMessage, error)

	AuthorizeCacheSecurityGroupIngress(*elasticache.AuthorizeCacheSecurityGroupIngressInput) (*elasticache.AuthorizeCacheSecurityGroupIngressOutput, error)

	CopySnapshot(*elasticache.CopySnapshotInput) (*elasticache.CopySnapshotOutput, error)

	CreateCacheCluster(*elasticache.CreateCacheClusterInput) (*elasticache.CreateCacheClusterOutput, error)

	CreateCacheParameterGroup(*elasticache.CreateCacheParameterGroupInput) (*elasticache.CreateCacheParameterGroupOutput, error)

	CreateCacheSecurityGroup(*elasticache.CreateCacheSecurityGroupInput) (*elasticache.CreateCacheSecurityGroupOutput, error)

	CreateCacheSubnetGroup(*elasticache.CreateCacheSubnetGroupInput) (*elasticache.CreateCacheSubnetGroupOutput, error)

	CreateReplicationGroup(*elasticache.CreateReplicationGroupInput) (*elasticache.CreateReplicationGroupOutput, error)

	CreateSnapshot(*elasticache.CreateSnapshotInput) (*elasticache.CreateSnapshotOutput, error)

	DeleteCacheCluster(*elasticache.DeleteCacheClusterInput) (*elasticache.DeleteCacheClusterOutput, error)

	DeleteCacheParameterGroup(*elasticache.DeleteCacheParameterGroupInput) (*elasticache.DeleteCacheParameterGroupOutput, error)

	DeleteCacheSecurityGroup(*elasticache.DeleteCacheSecurityGroupInput) (*elasticache.DeleteCacheSecurityGroupOutput, error)

	DeleteCacheSubnetGroup(*elasticache.DeleteCacheSubnetGroupInput) (*elasticache.DeleteCacheSubnetGroupOutput, error)

	DeleteReplicationGroup(*elasticache.DeleteReplicationGroupInput) (*elasticache.DeleteReplicationGroupOutput, error)

	DeleteSnapshot(*elasticache.DeleteSnapshotInput) (*elasticache.DeleteSnapshotOutput, error)

	DescribeCacheClusters(*elasticache.DescribeCacheClustersInput) (*elasticache.DescribeCacheClustersOutput, error)

	DescribeCacheEngineVersions(*elasticache.DescribeCacheEngineVersionsInput) (*elasticache.DescribeCacheEngineVersionsOutput, error)

	DescribeCacheParameterGroups(*elasticache.DescribeCacheParameterGroupsInput) (*elasticache.DescribeCacheParameterGroupsOutput, error)

	DescribeCacheParameters(*elasticache.DescribeCacheParametersInput) (*elasticache.DescribeCacheParametersOutput, error)

	DescribeCacheSecurityGroups(*elasticache.DescribeCacheSecurityGroupsInput) (*elasticache.DescribeCacheSecurityGroupsOutput, error)

	DescribeCacheSubnetGroups(*elasticache.DescribeCacheSubnetGroupsInput) (*elasticache.DescribeCacheSubnetGroupsOutput, error)

	DescribeEngineDefaultParameters(*elasticache.DescribeEngineDefaultParametersInput) (*elasticache.DescribeEngineDefaultParametersOutput, error)

	DescribeEvents(*elasticache.DescribeEventsInput) (*elasticache.DescribeEventsOutput, error)

	DescribeReplicationGroups(*elasticache.DescribeReplicationGroupsInput) (*elasticache.DescribeReplicationGroupsOutput, error)

	DescribeReservedCacheNodes(*elasticache.DescribeReservedCacheNodesInput) (*elasticache.DescribeReservedCacheNodesOutput, error)

	DescribeReservedCacheNodesOfferings(*elasticache.DescribeReservedCacheNodesOfferingsInput) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error)

	DescribeSnapshots(*elasticache.DescribeSnapshotsInput) (*elasticache.DescribeSnapshotsOutput, error)

	ListTagsForResource(*elasticache.ListTagsForResourceInput) (*elasticache.TagListMessage, error)

	ModifyCacheCluster(*elasticache.ModifyCacheClusterInput) (*elasticache.ModifyCacheClusterOutput, error)

	ModifyCacheParameterGroup(*elasticache.ModifyCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error)

	ModifyCacheSubnetGroup(*elasticache.ModifyCacheSubnetGroupInput) (*elasticache.ModifyCacheSubnetGroupOutput, error)

	ModifyReplicationGroup(*elasticache.ModifyReplicationGroupInput) (*elasticache.ModifyReplicationGroupOutput, error)

	PurchaseReservedCacheNodesOffering(*elasticache.PurchaseReservedCacheNodesOfferingInput) (*elasticache.PurchaseReservedCacheNodesOfferingOutput, error)

	RebootCacheCluster(*elasticache.RebootCacheClusterInput) (*elasticache.RebootCacheClusterOutput, error)

	RemoveTagsFromResource(*elasticache.RemoveTagsFromResourceInput) (*elasticache.TagListMessage, error)

	ResetCacheParameterGroup(*elasticache.ResetCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error)

	RevokeCacheSecurityGroupIngress(*elasticache.RevokeCacheSecurityGroupIngressInput) (*elasticache.RevokeCacheSecurityGroupIngressOutput, error)
}