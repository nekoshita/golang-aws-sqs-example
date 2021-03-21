resource "aws_iam_policy_attachment" "my_queue" {
  name       = "my_queue"
  policy_arn = aws_iam_policy.my_queue.arn
  users      = [aws_iam_user.my_queue.name]
}
