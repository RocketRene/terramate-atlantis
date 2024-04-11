resource "aws_s3_bucket" "bucket" {
  bucket = "terramate-rene-tyree994-movies"

  acl = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
