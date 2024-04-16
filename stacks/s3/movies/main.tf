resource "aws_s3_bucket" "bucket" {
  bucket = "terramate-rene-aidan539-movies"


  acl = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
