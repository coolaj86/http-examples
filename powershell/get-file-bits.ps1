# http://superuser.com/questions/362152/native-alternative-to-wget-in-windows-powershell
# This probably won't work for you

function HttpGet-BitsDownload()
{
  params (
    [string] $target = "http://foobar3000.com/echo" #$null
  )

  Import-Module BitsTransfer
  Start-BitsTransfer -source "$target"
}

HttpGet-BitsDownload "http://foobar3000.com/echo"
