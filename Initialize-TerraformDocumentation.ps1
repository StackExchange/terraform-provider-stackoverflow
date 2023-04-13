param(
    [string] $Path,
    [string] $OutputPath,
    [string[]] $Tags,
    [string] $FilterResource
)

$ResourcesJson = "{"
Get-ChildItem -Path $Path -Filter "*.md" | ForEach-Object {
    $Resource = @{
        stackoverflow_article = @{
            "$($_.BaseName -Replace $Pattern, '_')" = @{
                article_type = "knowledge-article"
                title = $_.BaseName
                body_markdown = "`${file(`"$($_.FullName -replace "\\", "/")`")}"
                tags = $Tags
                filter = $FilterResource
            }
        }
    }
    $ResourcesJson += "`"resource`": " + (ConvertTo-Json -Depth 10 -InputObject $Resource) + ","
}
$ResourcesJson = $ResourcesJson.TrimEnd(',')
$ResourcesJson += "}"

$ResourcesJson | Out-File -FilePath (Join-Path -Path $OutputPath -ChildPath "$(Split-Path -Path $Path -Leaf).tf.json")
