resource "aws_s3_bucket" "bucket" {
  bucket = "terramate-rene-alysha673-movies"

  acl = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
