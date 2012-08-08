function Execute-HTTPGetCommand()
{
  param(
    [string]$target = "http://foobar3000.com/echo" #$null
  )

  $webRequest = [System.Net.WebRequest]::Create($target)
  $webRequest.ServicePoint.Expect100Continue = $false
  $webRequest.Method = "Get"
  [System.Net.WebResponse]$resp = $webRequest.GetResponse()
  $rs = $resp.GetResponseStream()
  [System.IO.StreamReader]$sr = New-Object System.IO.StreamReader -argumentList $rs
  [string]$results = $sr.ReadToEnd()
  return $results
}

Execute-HTTPGetCommand "http://foobar3000.com/echo"
