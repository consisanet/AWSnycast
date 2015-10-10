package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"testing"
)

var (
	rtb1 = ec2.RouteTable{
		RouteTableId: aws.String("rtb-f0ea3b95"),
		Routes: []*ec2.Route{
			&ec2.Route{
				DestinationCidrBlock: aws.String("172.17.16.0/22"),
				GatewayId:            aws.String("local"),
				Origin:               aws.String("CreateRouteTable"),
				State:                aws.String("active"),
			},
		},
		Tags: []*ec2.Tag{
			&ec2.Tag{
				Key:   aws.String("Name"),
				Value: aws.String("uswest1 devb private insecure"),
			}},
		VpcId: aws.String("vpc-9496cffc"),
	}

	rtb2 = ec2.RouteTable{
		Associations: []*ec2.RouteTableAssociation{
			&ec2.RouteTableAssociation{
				Main: aws.Bool(true),
				RouteTableAssociationId: aws.String("rtbassoc-b1f025d4"),
				RouteTableId:            aws.String("rtb-9696cffe"),
			},
			&ec2.RouteTableAssociation{
				Main: aws.Bool(false),
				RouteTableAssociationId: aws.String("rtbassoc-85c1cbe7"),
				RouteTableId:            aws.String("rtb-9696cffe"),
				SubnetId:                aws.String("subnet-16b0e97e"),
			},
			&ec2.RouteTableAssociation{
				Main: aws.Bool(false),
				RouteTableAssociationId: aws.String("rtbassoc-ba8573df"),
				RouteTableId:            aws.String("rtb-9696cffe"),
				SubnetId:                aws.String("subnet-3fb0e957"),
			},
			&ec2.RouteTableAssociation{
				Main: aws.Bool(false),
				RouteTableAssociationId: aws.String("rtbassoc-84c1cbe6"),
				RouteTableId:            aws.String("rtb-9696cffe"),
				SubnetId:                aws.String("subnet-28b0e940"),
			},
			&ec2.RouteTableAssociation{
				Main: aws.Bool(false),
				RouteTableAssociationId: aws.String("rtbassoc-858573e0"),
				RouteTableId:            aws.String("rtb-9696cffe"),
				SubnetId:                aws.String("subnet-f3b0e99b"),
			},
		},
		PropagatingVgws: []*ec2.PropagatingVgw{
			&ec2.PropagatingVgw{
				GatewayId: aws.String("vgw-d2396a97"),
			},
		},
		RouteTableId: aws.String("rtb-9696cffe"),
		Routes: []*ec2.Route{
			&ec2.Route{
				DestinationCidrBlock: aws.String("10.55.35.43/32"),
				GatewayId:            aws.String("vgw-d2396a97"),
				Origin:               aws.String("CreateRoute"),
				State:                aws.String("active"),
			},
			&ec2.Route{
				DestinationCidrBlock: aws.String("172.17.16.0/22"),
				GatewayId:            aws.String("local"),
				Origin:               aws.String("CreateRouteTable"),
				State:                aws.String("active"),
			},
			&ec2.Route{
				DestinationCidrBlock: aws.String("10.0.0.0/8"),
				InstanceId:           aws.String("i-446b201b"),
				InstanceOwnerId:      aws.String("613514870339"),
				NetworkInterfaceId:   aws.String("eni-ea8a9cac"),
				Origin:               aws.String("CreateRoute"),
				State:                aws.String("active"),
			},
			&ec2.Route{
				DestinationCidrBlock: aws.String("0.0.0.0/0"),
				InstanceId:           aws.String("i-605bd2aa"),
				InstanceOwnerId:      aws.String("613514870339"),
				NetworkInterfaceId:   aws.String("eni-09472250"),
				Origin:               aws.String("CreateRoute"),
				State:                aws.String("active"),
			},
			&ec2.Route{
				DestinationCidrBlock: aws.String("10.0.0.0/8"),
				GatewayId:            aws.String("vgw-d2396a97"),
				Origin:               aws.String("EnableVgwRoutePropagation"),
				State:                aws.String("active"),
			},
		},
		Tags: []*ec2.Tag{
			&ec2.Tag{
				Key:   aws.String("Name"),
				Value: aws.String("uswest1 devb private"),
			}},
		VpcId: aws.String("vpc-9496cffc"),
	}

	rtb3 = ec2.RouteTable{
		Associations: []*ec2.RouteTableAssociation{
			&ec2.RouteTableAssociation{
				Main: aws.Bool(false),
				RouteTableAssociationId: aws.String("rtbassoc-818573e4"),
				RouteTableId:            aws.String("rtb-019cab69"),
				SubnetId:                aws.String("subnet-37b0e95f"),
			},
			&ec2.RouteTableAssociation{
				Main: aws.Bool(false),
				RouteTableAssociationId: aws.String("rtbassoc-fd9cab95"),
				RouteTableId:            aws.String("rtb-019cab69"),
				SubnetId:                aws.String("subnet-44b0e92c"),
			},
		},
		PropagatingVgws: []*ec2.PropagatingVgw{
			&ec2.PropagatingVgw{
				GatewayId: aws.String("vgw-d2396a97"),
			},
		},
		RouteTableId: aws.String("rtb-019cab69"),
		Routes: []*ec2.Route{
			&ec2.Route{
				DestinationCidrBlock: aws.String("10.55.35.43/32"),
				GatewayId:            aws.String("vgw-d2396a97"),
				Origin:               aws.String("CreateRoute"),
				State:                aws.String("active"),
			},
			&ec2.Route{
				DestinationCidrBlock: aws.String("172.17.16.0/22"),
				GatewayId:            aws.String("local"),
				Origin:               aws.String("CreateRouteTable"),
				State:                aws.String("active"),
			},
			&ec2.Route{
				DestinationCidrBlock: aws.String("10.0.0.0/8"),
				InstanceId:           aws.String("i-446b201b"),
				InstanceOwnerId:      aws.String("613514870339"),
				NetworkInterfaceId:   aws.String("eni-ea8a9cac"),
				Origin:               aws.String("CreateRoute"),
				State:                aws.String("active"),
			},
			&ec2.Route{
				DestinationCidrBlock: aws.String("0.0.0.0/0"),
				GatewayId:            aws.String("igw-9ab1e8f2"),
				Origin:               aws.String("CreateRoute"),
				State:                aws.String("active"),
			},
			&ec2.Route{
				DestinationCidrBlock: aws.String("10.0.0.0/8"),
				GatewayId:            aws.String("vgw-d2396a97"),
				Origin:               aws.String("EnableVgwRoutePropagation"),
				State:                aws.String("active"),
			},
		},
		Tags: []*ec2.Tag{
			&ec2.Tag{
				Key:   aws.String("Name"),
				Value: aws.String("uswest1 devb public"),
			},
		},
		VpcId: aws.String("vpc-9496cffc"),
	}

	rtb4 = ec2.RouteTable{
		RouteTableId: aws.String("rtb-f1ea3b94"),
		Routes: []*ec2.Route{
			&ec2.Route{
				DestinationCidrBlock: aws.String("172.17.16.0/22"),
				GatewayId:            aws.String("local"),
				Origin:               aws.String("CreateRouteTable"),
				State:                aws.String("active"),
			},
			&ec2.Route{
				DestinationCidrBlock: aws.String("0.0.0.0/0"),
				GatewayId:            aws.String("igw-9ab1e8f2"),
				Origin:               aws.String("CreateRoute"),
				State:                aws.String("active"),
			},
		},
		Tags: []*ec2.Tag{
			&ec2.Tag{
				Key:   aws.String("Name"),
				Value: aws.String("uswest1 devb public insecure"),
			},
		},
		VpcId: aws.String("vpc-9496cffc"),
	}
)

type FakeRouteTableFetcher struct {
	Error  error
	Routes []*ec2.RouteTable
}

func (r FakeRouteTableFetcher) GetRouteTables() ([]*ec2.RouteTable, error) {
	return r.Routes, r.Error
}

func TestFakeFetcher(t *testing.T) {
	var f RouteTableFetcher
	f = FakeRouteTableFetcher{
		Routes: []*ec2.RouteTable{&rtb1},
	}
	rtb, err := f.GetRouteTables()
	if err != nil {
		t.Fail()
	}
	if len(rtb) != 1 {
		t.Fail()
	}
	if rtb[0] != &rtb1 {
		t.Fail()
	}
}

func TestRouteTableFilterAlways(t *testing.T) {
	f := RouteTableFilterAlways{}
	if f.Keep(&rtb1) {
		t.Fail()
	}
}

func TestRouteTableFilterNever(t *testing.T) {
	f := RouteTableFilterNever{}
	if !f.Keep(&rtb1) {
		t.Fail()
	}
}

func RouteTableFilterAndTwoNever(t *testing.T) {
	f := RouteTableFilterAnd{
		RouteTableFilters: []RouteTableFilter{
			RouteTableFilterNever{},
			RouteTableFilterNever{},
		},
	}
	if !f.Keep(&rtb1) {
		t.Fail()
	}
}

func RouteTableFilterAndOneNever(t *testing.T) {
	f := RouteTableFilterAnd{
		RouteTableFilters: []RouteTableFilter{
			RouteTableFilterNever{},
			RouteTableFilterAlways{},
		},
	}
	if f.Keep(&rtb1) {
		t.Fail()
	}
}

func RouteTableFilterOrOneNever(t *testing.T) {
	f := RouteTableFilterAnd{
		RouteTableFilters: []RouteTableFilter{
			RouteTableFilterNever{},
			RouteTableFilterAlways{},
		},
	}
	if !f.Keep(&rtb1) {
		t.Fail()
	}
}

func RouteTableFilterOrOneNever2(t *testing.T) {
	f := RouteTableFilterAnd{
		RouteTableFilters: []RouteTableFilter{
			RouteTableFilterAlways{},
			RouteTableFilterNever{},
		},
	}
	if !f.Keep(&rtb1) {
		t.Fail()
	}
}

func RouteTableFilterOrAlways(t *testing.T) {
	f := RouteTableFilterAnd{
		RouteTableFilters: []RouteTableFilter{
			RouteTableFilterAlways{},
			RouteTableFilterAlways{},
		},
	}
	if f.Keep(&rtb1) {
		t.Fail()
	}
}

func TestFilterRouteTables(t *testing.T) {
	rtb := FilterRouteTables(RouteTableFilterNever{}, []*ec2.RouteTable{&rtb1})
	if len(rtb) != 1 {
		t.Fail()
	}
	if rtb[0] != &rtb1 {
		t.Fail()
	}
}
