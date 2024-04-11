resource "aws_s3_bucket" "bucket" {
  bucket = "terramate-rene-pierce822-movies"

  acl = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
