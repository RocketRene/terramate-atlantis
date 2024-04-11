resource "aws_s3_bucket" "bucket" {
  bucket = "terramate-rene-linnie817-movies"

  acl = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
