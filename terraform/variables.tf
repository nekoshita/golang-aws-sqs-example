variable "regions" {
  type = map(string)
  default = {
    "tokyo" = "ap-northeast-1"
  }
}

variable "zone" {
  type = map(string)
  default = {
    "tokyo-a" = "ap-northeast-1a"
    "tokyo-c" = "ap-northeast-1c"
  }
}
