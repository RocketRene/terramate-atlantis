resource "aws_s3_bucket" "bucket" {
  bucket = "terramate-rene-alessandra113-movies"


  acl = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
