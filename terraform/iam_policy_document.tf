data "aws_iam_policy_document" "my_queue" {
  statement {
    actions = [
      "sqs:SendMessage",
      "sqs:ReceiveMessage",
      "sqs:DeleteMessage",
    ]

    resources = [aws_sqs_queue.my_queue.arn]
  }
}
