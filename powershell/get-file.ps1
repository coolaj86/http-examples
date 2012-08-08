# http://superuser.com/questions/362152/native-alternative-to-wget-in-windows-powershell

function HttpGet-Download()
{
  params (
    [string] $target = "http://foobar3000.com/echo.json"
    [string] $path = "./tmp-download.bin"
  )

  (new-object System.Net.WebClient).DownloadFile( '$url, $path)
}

# Same as
#$client = new-object System.Net.WebClient
#$client.DownloadFile( $url, $path )
