class UrlencodeDecode < Formula
  desc "URL-encode and decode text from arguments or standard input"
  homepage "https://github.com/argium/go-urlencode-urldecode"
  head "https://github.com/argium/go-urlencode-urldecode.git", branch: "main"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(output: bin/"urlencode"), "./cmd/urlencode"
    system "go", "build", *std_go_args(output: bin/"urldecode"), "./cmd/urldecode"
  end
end
