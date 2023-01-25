// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package symbiosis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Kubernetes clusters service accounts.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/symbiosis-cloud/pulumi-symbiosis/sdk/go/symbiosis"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := symbiosis.NewClusterServiceAccount(ctx, "example", &symbiosis.ClusterServiceAccountArgs{
//				ClusterName: pulumi.Any(symbiosis_cluster.Example.Name),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ClusterServiceAccount struct {
	pulumi.CustomResourceState

	// Cluster CA certificate
	ClusterCaCertificate pulumi.StringOutput `pulumi:"clusterCaCertificate"`
	// Cluster name. Changing the name forces re-creation.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Service account token
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewClusterServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewClusterServiceAccount(ctx *pulumi.Context,
	name string, args *ClusterServiceAccountArgs, opts ...pulumi.ResourceOption) (*ClusterServiceAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clusterCaCertificate",
		"token",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource ClusterServiceAccount
	err := ctx.RegisterResource("symbiosis:index/clusterServiceAccount:ClusterServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterServiceAccount gets an existing ClusterServiceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterServiceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterServiceAccountState, opts ...pulumi.ResourceOption) (*ClusterServiceAccount, error) {
	var resource ClusterServiceAccount
	err := ctx.ReadResource("symbiosis:index/clusterServiceAccount:ClusterServiceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterServiceAccount resources.
type clusterServiceAccountState struct {
	// Cluster CA certificate
	ClusterCaCertificate *string `pulumi:"clusterCaCertificate"`
	// Cluster name. Changing the name forces re-creation.
	ClusterName *string `pulumi:"clusterName"`
	// Service account token
	Token *string `pulumi:"token"`
}

type ClusterServiceAccountState struct {
	// Cluster CA certificate
	ClusterCaCertificate pulumi.StringPtrInput
	// Cluster name. Changing the name forces re-creation.
	ClusterName pulumi.StringPtrInput
	// Service account token
	Token pulumi.StringPtrInput
}

func (ClusterServiceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterServiceAccountState)(nil)).Elem()
}

type clusterServiceAccountArgs struct {
	// Cluster name. Changing the name forces re-creation.
	ClusterName string `pulumi:"clusterName"`
}

// The set of arguments for constructing a ClusterServiceAccount resource.
type ClusterServiceAccountArgs struct {
	// Cluster name. Changing the name forces re-creation.
	ClusterName pulumi.StringInput
}

func (ClusterServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterServiceAccountArgs)(nil)).Elem()
}

type ClusterServiceAccountInput interface {
	pulumi.Input

	ToClusterServiceAccountOutput() ClusterServiceAccountOutput
	ToClusterServiceAccountOutputWithContext(ctx context.Context) ClusterServiceAccountOutput
}

func (*ClusterServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterServiceAccount)(nil)).Elem()
}

func (i *ClusterServiceAccount) ToClusterServiceAccountOutput() ClusterServiceAccountOutput {
	return i.ToClusterServiceAccountOutputWithContext(context.Background())
}

func (i *ClusterServiceAccount) ToClusterServiceAccountOutputWithContext(ctx context.Context) ClusterServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterServiceAccountOutput)
}

// ClusterServiceAccountArrayInput is an input type that accepts ClusterServiceAccountArray and ClusterServiceAccountArrayOutput values.
// You can construct a concrete instance of `ClusterServiceAccountArrayInput` via:
//
//	ClusterServiceAccountArray{ ClusterServiceAccountArgs{...} }
type ClusterServiceAccountArrayInput interface {
	pulumi.Input

	ToClusterServiceAccountArrayOutput() ClusterServiceAccountArrayOutput
	ToClusterServiceAccountArrayOutputWithContext(context.Context) ClusterServiceAccountArrayOutput
}

type ClusterServiceAccountArray []ClusterServiceAccountInput

func (ClusterServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterServiceAccount)(nil)).Elem()
}

func (i ClusterServiceAccountArray) ToClusterServiceAccountArrayOutput() ClusterServiceAccountArrayOutput {
	return i.ToClusterServiceAccountArrayOutputWithContext(context.Background())
}

func (i ClusterServiceAccountArray) ToClusterServiceAccountArrayOutputWithContext(ctx context.Context) ClusterServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterServiceAccountArrayOutput)
}

// ClusterServiceAccountMapInput is an input type that accepts ClusterServiceAccountMap and ClusterServiceAccountMapOutput values.
// You can construct a concrete instance of `ClusterServiceAccountMapInput` via:
//
//	ClusterServiceAccountMap{ "key": ClusterServiceAccountArgs{...} }
type ClusterServiceAccountMapInput interface {
	pulumi.Input

	ToClusterServiceAccountMapOutput() ClusterServiceAccountMapOutput
	ToClusterServiceAccountMapOutputWithContext(context.Context) ClusterServiceAccountMapOutput
}

type ClusterServiceAccountMap map[string]ClusterServiceAccountInput

func (ClusterServiceAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterServiceAccount)(nil)).Elem()
}

func (i ClusterServiceAccountMap) ToClusterServiceAccountMapOutput() ClusterServiceAccountMapOutput {
	return i.ToClusterServiceAccountMapOutputWithContext(context.Background())
}

func (i ClusterServiceAccountMap) ToClusterServiceAccountMapOutputWithContext(ctx context.Context) ClusterServiceAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterServiceAccountMapOutput)
}

type ClusterServiceAccountOutput struct{ *pulumi.OutputState }

func (ClusterServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterServiceAccount)(nil)).Elem()
}

func (o ClusterServiceAccountOutput) ToClusterServiceAccountOutput() ClusterServiceAccountOutput {
	return o
}

func (o ClusterServiceAccountOutput) ToClusterServiceAccountOutputWithContext(ctx context.Context) ClusterServiceAccountOutput {
	return o
}

// Cluster CA certificate
func (o ClusterServiceAccountOutput) ClusterCaCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterServiceAccount) pulumi.StringOutput { return v.ClusterCaCertificate }).(pulumi.StringOutput)
}

// Cluster name. Changing the name forces re-creation.
func (o ClusterServiceAccountOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterServiceAccount) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// Service account token
func (o ClusterServiceAccountOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterServiceAccount) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

type ClusterServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (ClusterServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterServiceAccount)(nil)).Elem()
}

func (o ClusterServiceAccountArrayOutput) ToClusterServiceAccountArrayOutput() ClusterServiceAccountArrayOutput {
	return o
}

func (o ClusterServiceAccountArrayOutput) ToClusterServiceAccountArrayOutputWithContext(ctx context.Context) ClusterServiceAccountArrayOutput {
	return o
}

func (o ClusterServiceAccountArrayOutput) Index(i pulumi.IntInput) ClusterServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterServiceAccount {
		return vs[0].([]*ClusterServiceAccount)[vs[1].(int)]
	}).(ClusterServiceAccountOutput)
}

type ClusterServiceAccountMapOutput struct{ *pulumi.OutputState }

func (ClusterServiceAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterServiceAccount)(nil)).Elem()
}

func (o ClusterServiceAccountMapOutput) ToClusterServiceAccountMapOutput() ClusterServiceAccountMapOutput {
	return o
}

func (o ClusterServiceAccountMapOutput) ToClusterServiceAccountMapOutputWithContext(ctx context.Context) ClusterServiceAccountMapOutput {
	return o
}

func (o ClusterServiceAccountMapOutput) MapIndex(k pulumi.StringInput) ClusterServiceAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterServiceAccount {
		return vs[0].(map[string]*ClusterServiceAccount)[vs[1].(string)]
	}).(ClusterServiceAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterServiceAccountInput)(nil)).Elem(), &ClusterServiceAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterServiceAccountArrayInput)(nil)).Elem(), ClusterServiceAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterServiceAccountMapInput)(nil)).Elem(), ClusterServiceAccountMap{})
	pulumi.RegisterOutputType(ClusterServiceAccountOutput{})
	pulumi.RegisterOutputType(ClusterServiceAccountArrayOutput{})
	pulumi.RegisterOutputType(ClusterServiceAccountMapOutput{})
}
