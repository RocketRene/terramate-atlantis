resource "aws_s3_bucket" "bucket" {
  bucket = "terramate-rene-tremayne479-movies"


  acl = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
