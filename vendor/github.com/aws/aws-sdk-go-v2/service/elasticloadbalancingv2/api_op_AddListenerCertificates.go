// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds the specified SSL server certificate to the certificate list for the
// specified HTTPS or TLS listener.
//
// If the certificate in already in the certificate list, the call is successful
// but the certificate is not added again.
//
// For more information, see [HTTPS listeners] in the Application Load Balancers Guide or [TLS listeners] in the
// Network Load Balancers Guide.
//
// [TLS listeners]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html
// [HTTPS listeners]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html
func (c *Client) AddListenerCertificates(ctx context.Context, params *AddListenerCertificatesInput, optFns ...func(*Options)) (*AddListenerCertificatesOutput, error) {
	if params == nil {
		params = &AddListenerCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddListenerCertificates", params, optFns, c.addOperationAddListenerCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddListenerCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddListenerCertificatesInput struct {

	// The certificate to add. You can specify one certificate per call. Set
	// CertificateArn to the certificate ARN but do not set IsDefault .
	//
	// This member is required.
	Certificates []types.Certificate

	// The Amazon Resource Name (ARN) of the listener.
	//
	// This member is required.
	ListenerArn *string

	noSmithyDocumentSerde
}

type AddListenerCertificatesOutput struct {

	// Information about the certificates in the certificate list.
	Certificates []types.Certificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddListenerCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddListenerCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddListenerCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddListenerCertificates"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpAddListenerCertificatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddListenerCertificates(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAddListenerCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddListenerCertificates",
	}
}