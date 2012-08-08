# http://www.techmumbojumblog.com/?p=306
function ConvertTo-Base64($string) {
   $bytes  = [System.Text.Encoding]::UTF8.GetBytes($string);
   $encoded = [System.Convert]::ToBase64String($bytes); 

   return $encoded;
}

function ConvertFrom-Base64($string) {
   $bytes  = [System.Convert]::FromBase64String($string);
   $decoded = [System.Text.Encoding]::UTF8.GetString($bytes); 

   return $decoded;
}

# http://stackoverflow.com/questions/8919414/powershell-http-post-rest-api-basic-authentication
function Execute-HTTPPostCommand() {
    param(
        [string] $target = $null
        [string] $body = $null
    )

    $username = "user"
    $password = "pass"
    $auth = "Basic " + ConvertTo-Base64($username + ":" $password)

    $webRequest = [System.Net.WebRequest]::Create($target)
    $webRequest.ContentType = "application/json"
    $BodyStr = [System.Text.Encoding]::UTF8.GetBytes($body)
    $webrequest.ContentLength = $BodyStr.Length
    $webRequest.ServicePoint.Expect100Continue = $false
    #$webRequest.Headers.Add("Authorization", $auth);

    $webRequest.PreAuthenticate = $true
    $webRequest.Method = "POST"

    $requestStream = $webRequest.GetRequestStream()
    $requestStream.Write($BodyStr, 0,$BodyStr.length)
    $requestStream.Close()

    [System.Net.WebResponse] $resp = $webRequest.GetResponse()
    $rs = $resp.GetResponseStream()
    [System.IO.StreamReader] $sr = New-Object System.IO.StreamReader -argumentList $rs
    [string] $results = $sr.ReadToEnd()

    return $results
}

Execute-HTTPPostCommand "http://foobar3000.com/echo/" '{"foo":"boo"}'
