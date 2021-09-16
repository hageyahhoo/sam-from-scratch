# sam-from-scratch

This project aims to learn how to create an [AWS Lambda function](https://aws.amazon.com/lambda/) which interacts with [Amazon DynamoDB](https://aws.amazon.com/dynamodb/).
<br><br><br>



# How to Use This Project

## 1. DynamoDB
1. Create a table on DynamoDB by running the following command.
    ```
    aws dynamodb create-table --cli-input-json file://./dynamodb/table.json
    ```
1. Input data into the table by running the following command.
    ```
    aws dynamodb batch-write-item --request-items file://./dynamodb/data.json
    ```
1. You can scan data on the table by running the following command.
    ```
    aws dynamodb scan --table-name <TABLE_NAME>
    ```
<br>


## 2. Docker Image
1. Create a repository on [Amazon ECR](https://aws.amazon.com/ecr/) by running the following command.
    ```
    aws ecr create-repository --repository-name samscratch/service
    ```
1. Login to ECR by running the following command.
    ```
    aws ecr get-login-password --region <YOUR_AWS_REGION> | docker login --username AWS --password-stdin <YOUR_ACCOUNT_ID>.dkr.ecr.<YOUR_AWS_REGION>.amazonaws.com
    ```
1. Build a Docker image by running the following command.
    ```
    docker build -t samscratch/service .
    ```
1. Add a tag to the Docker image built by running the following command.
    ```
	docker tag samscratch/service:latest <YOUR_ACCOUNT_ID>.dkr.ecr.<YOUR_AWS_REGION>.amazonaws.com/samscratch/service:latest
    ```
1. Push the Docker image to ECR by running the following command.
    ```
	docker push <YOUR_ACCOUNT_ID>.dkr.ecr.<YOUR_AWS_REGION>.amazonaws.com/samscratch/service:latest
    ```
<br>


## 3. Lambda
1. Create your Lambda function and related IAM role by running the following command.
    ```
    aws cloudformation create-stack --stack-name samScratchStack --capabilities CAPABILITY_NAMED_IAM --template-body file://./cfn/lambda.yaml
    ```
1. Invoke your Lambda function by running the following command.
    ```
    aws lambda invoke --function-name samScratch out --log-type Tail --query 'LogResult' --output text | base64 -d
    ```
    - You can add `--debug` option to know more detail of the invocation.
<br><br>


## 4. API Gateway
1. Create your [API Gateway](https://aws.amazon.com/api-gateway/) by following on [this page](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-create-api-as-simple-proxy-for-lambda.html#api-gateway-create-api-as-simple-proxy-for-lambda-build).
1. Deploy and test your API Gateway by following on [this page](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-create-api-as-simple-proxy-for-lambda.html#api-gateway-create-api-as-simple-proxy-for-lambda-test)
<br><br><br>



# How to Use This Project from Fargate
You can use [AWS Fargate](https://aws.amazon.com/fargate/) instead of Lambda function.
<br><br>


## 1. DynamoDB
The same as Lambda.
<br><br>


## 2. Docker Image
1. Create a repository on ECR by running the following command.
    ```
    aws ecr create-repository --repository-name samscratch/fargate
    ```
1. Login to ECR by running the following command.
    ```
    aws ecr get-login-password --region <YOUR_AWS_REGION> | docker login --username AWS --password-stdin <YOUR_ACCOUNT_ID>.dkr.ecr.<YOUR_AWS_REGION>.amazonaws.com
    ```
1. Build a Docker image by running the following command.
    ```
    docker build --file ./FargateDockerfile -t samscratch/fargate .
    ```
1. Add a tag to the Docker image built by running the following command.
    ```
	docker tag samscratch/fargate:latest <YOUR_ACCOUNT_ID>.dkr.ecr.<YOUR_AWS_REGION>.amazonaws.com/samscratch/fargate:latest
    ```
1. Push the Docker image to ECR by running the following command.
    ```
	docker push <YOUR_ACCOUNT_ID>.dkr.ecr.<YOUR_AWS_REGION>.amazonaws.com/samscratch/fargate:latest
    ```
<br>


## 3. Fargate
1. Set up Fargate and related IAM role by running the following command.
    ```
    aws cloudformation create-stack --stack-name samScratchFargate --capabilities CAPABILITY_NAMED_IAM --template-body file://./cfn/fargate.yaml
    ```
1. Check `DNS name` of ELB.
1. Invoke your Lambda function by running the following command.
    ```
    curl http://<DNS name of ELB>/samScratch
    ```
<br><br><br>



# How to Use This Project with GraphQL
You can use [GraphQL](https://graphql.org/) API with [AWS AppSync](https://aws.amazon.com/appsync/) instead of `API Gateway`.
<br><br>


## 1. DynamoDB, Docker Image, and Lambda
The same as Lambda.
<br><br>


## 2. GraphQL
1. Create S3 Bucket and upload `schema.graphql` to that.
    1. Create S3 Bucket by running the following command.
        ```
        aws s3 mb s3://sam-scratch-graphql
        ```
        - BucketName: `sam-scratch-graphql`
    1. Upload `schema.graphql` to S3 Bucket created by running the following command.
        ```
        aws s3 cp ./graphql/schema.graphql s3://sam-scratch-graphql/schema.graphql
        ```
    1. Apply the policy file to S3 Bucket by running the following command.
        ```
        aws s3api put-bucket-policy --bucket sam-scratch-graphql --policy file://./graphql/bucket_policy.json
        ```
1. Create your `Schema`, `Data Source`, `Resolver` and related IAM role by running the following command.
    ```
    aws cloudformation create-stack --stack-name samScratchGraphQL --capabilities CAPABILITY_NAMED_IAM --template-body file://./cfn/appsync.yaml
    ```
1. Access to AppSync Console and get `API URL` and `API KEY`.
1. Invoke your GraphQL API by running the following commands.
    ```
    curl -H "Content-Type:application/graphql" \
    -H "x-api-key:<API KEY>" \
    -d '{ "query": "query allServants { allServants { ServantId Name Class } }"}' \
    <API URL>
    ```
    or
    ```
    curl -H "Content-Type:application/graphql" \
    -H "x-api-key:<API KEY>" \
    -d '{ "query": "query singleServant { singleServant(ServantId:\"23\") { ServantId Name Class } }"}' \
    <API URL>
    ```
    1. Specify `x-api-key:<API KEY>` as HTTP Header.
        - https://docs.aws.amazon.com/appsync/latest/devguide/security-authz.html#api-key-authorization
    1. Specify `Content-Type:application/graphql` as HTTP Header.
    1. Specify GraphQL query like `"query": "query <QUERY_NAME> { QUERY_NAME { ITEM_NAME } }`.
<br><br>
