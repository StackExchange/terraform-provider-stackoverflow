$terraformVersion = (Get-Content -Path '.terraform-version').Trim()

if ($null -ne $env:CI) {
    Write-Output -InputObject "value=$terraformVersion" >> $env:GITHUB_OUTPUT
}
