resource "aws_s3_bucket" "bucket" {
  bucket = "terramate-rene-einar830-movies"


  acl = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
