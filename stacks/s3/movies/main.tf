resource "aws_s3_bucket" "bucket" {
  bucket = "terramate-rene-avery772-movies"


  acl = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
