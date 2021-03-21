resource "aws_sqs_queue" "my_queue" {
  name                        = "my_queue.fifo"
  fifo_queue                  = true
  content_based_deduplication = true
  redrive_policy              = "{\"deadLetterTargetArn\":\"${aws_sqs_queue.deadletter.arn}\",\"maxReceiveCount\":3}"
}

resource "aws_sqs_queue" "deadletter" {
  name                      = "deadletter.fifo"
  message_retention_seconds = 1209600
  fifo_queue                = true
}
