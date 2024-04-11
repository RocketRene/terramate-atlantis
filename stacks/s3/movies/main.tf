resource "aws_s3_bucket" "bucket" {
  bucket = "terramate-rene-hector437-movies"

  acl = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
