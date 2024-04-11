resource "aws_s3_bucket" "bucket" {
  bucket = "terramate-rene-lilian376-movies"


  acl = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
