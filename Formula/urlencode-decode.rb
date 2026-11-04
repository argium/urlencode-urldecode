class UrlencodeDecode < Formula
  desc "URL-encode and decode text from arguments or standard input"
  homepage "https://github.com/argium/urlencode-urldecode"
  url "https://github.com/argium/urlencode-urldecode/releases/download/v0.1.0/urlencode-decode-v0.1.0.tar.gz"
  sha256 "203a0e32f1d08f786c7fcbc314d5cb4548031788b11d630bfa41f8e5124b0b9d"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(output: bin/"urlencode"), "./cmd/urlencode"
    system "go", "build", *std_go_args(output: bin/"urldecode"), "./cmd/urldecode"
  end
end
