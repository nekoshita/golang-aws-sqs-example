resource "aws_iam_policy" "my_queue" {
  name   = "my_queue"
  policy = data.aws_iam_policy_document.my_queue.json
}
