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
<br><br>
