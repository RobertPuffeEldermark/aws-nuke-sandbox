################## 
## Nuke Box
################## 


## Install git

sudo yum install git

## configure cli credentials

aws configure

# edit the files to use temporary credentials

cd .aws
nano credentials

# get temporary creds from AWS Identity center

## example 
[default]
aws_access_key_id=ASIAR6PEBF7AYRYPB4WC
aws_secret_access_key=6+jaADxUCUoNgS9OyOjW7kRbb2IrnYgDGfITxiut
aws_session_token=IQoJb3JpZ2luX2VjEOL//////////wEaCXVzLWVhc3QtMSJHMEUCIBI/yfOjKyOt+lBQq6j3mY+jB63EyLQ3JJMMKlPMoWEeAiEA4sgaoLhPLc5fKu>

# Connect to github [Title](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/managing-deploy-keys#deploy-keys)

# note update to something else eventually
git clone https://github.com/RobertPuffeEldermark/aws-nuke-sandbox.git

## Create IAM user with programmtic access
aws iam create-user --user-name aws-nuke
aws iam create-access-key --user-name aws-nuke
## Assign administrator permissions to the user
aws iam attach-user-policy --policy-arn arn:aws:iam:$ACCOUNTID:aws:policy/AdministratorAccess --user-name aws-nuke

## Change AWS Configuration to this user
aws configure
# use the access key created in the previous step. It is recommended to delete this key when done.

## Download aws-nuke
wget -c https://github.com/rebuy-de/aws-nuke/releases/download/v2.16.0/aws-nuke-v2.16.0-linux-amd64.tar.gz
## Extract the aws-nuke binary
tar -xvf aws-nuke-v2.16.0-linux-amd64.tar.gz
## Rename the extracted binary to aws-nuke
mv aws-nuke-v2.16.0-linux-amd64 aws-nuke
## Copy the extracted binary to your $PATH
sudo mv aws-nuke /usr/local/bin/aws-nuke
## Validate
aws-nuke -h


# change to the directory where the config.yml file is located

cd workspace\aws-nuke-sandbox\

# validate config file

nano config.yml

# run a test nuke

aws-nuke -c nuke-config.yml --profile aws_nuke

# If the test looks good, run the real nuke

add --no-dry-run to the previous command








###Copied from Medium

## Store AWS account ID in a variable
ACCOUNTID=$(aws sts get-caller-identity --query Account --output text)
## Create IAM user with programmtic access
aws iam create-user --user-name aws-nuke
aws iam create-access-key --user-name aws-nuke
## Assign administrator permissions to the user
aws iam attach-user-policy --policy-arn arn:aws:iam:$ACCOUNTID:aws:policy/AdministratorAccess --user-name aws-nuke


## Configure AWS profile
aws configure — profile aws_nuke
Enter your IAM ACCESS_KEY
Enter your IAM SECRET_KEY
Enter a default region E.g. us-east-2
Enter an output format E.g. json
## Configure AWS credential
cat .aws/credentials
cat .aws/config
aws sts get-caller-identity — profile aws_nuke
## Create an account alias name for your AWS account
aws iam create-account-alias — profile aws_nuke — account-alias dokeyode-aws-account


## Download aws-nuke
wget -c https://github.com/rebuy-de/aws-nuke/releases/download/v2.16.0/aws-nuke-v2.16.0-linux-amd64.tar.gz
## Extract the aws-nuke binary
tar -xvf aws-nuke-v2.16.0-linux-amd64.tar.gz
## Rename the extracted binary to aws-nuke
mv aws-nuke-v2.16.0-linux-amd64 aws-nuke
## Copy the extracted binary to your $PATH
sudo mv aws-nuke /usr/local/bin/aws-nuke
## Validate
aws-nuke -h
