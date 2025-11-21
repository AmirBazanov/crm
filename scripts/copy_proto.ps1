$SOURCE = "proto"
$DEST = "dist/proto"

Write-Host "Copying .proto files from $SOURCE to $DEST..."

# Ensure destination exists and is clean
if (Test-Path $DEST) {
    Remove-Item -Path $DEST -Recurse -Force
}
New-Item -ItemType Directory -Path $DEST -Force | Out-Null

# Get all .proto files
$files = Get-ChildItem -Path $SOURCE -Recurse -Filter "*.proto"

foreach ($file in $files) {
    # Calculate relative path
    $relativePath = $file.FullName.Substring((Resolve-Path $SOURCE).Path.Length + 1)
    
    # Construct destination path
    $destPath = Join-Path $DEST $relativePath
    $destDir = Split-Path $destPath -Parent

    # Create directory if it doesn't exist
    if (!(Test-Path $destDir)) {
        New-Item -ItemType Directory -Path $destDir -Force | Out-Null
    }

    # Copy file
    Copy-Item -Path $file.FullName -Destination $destPath -Force
    Write-Host "Copied: $relativePath"
}

# Copy gen directory if it exists
$genSource = Join-Path $SOURCE "gen"
$genDest = Join-Path $DEST "gen"

if (Test-Path $genSource) {
    Write-Host "Copying gen directory..."
    Copy-Item -Path $genSource -Destination $DEST -Recurse -Force
    Write-Host "Copied: gen directory"
}

Write-Host "Done."
