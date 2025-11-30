# Path to your proto directory (relative)
$ProtoDir = "proto"

# Output directory
$OutDir = "generated"

# Ensure output folder exists
if (!(Test-Path -Path $OutDir)) {
    New-Item -ItemType Directory -Force -Path $OutDir | Out-Null
}

# Get all .proto files as RELATIVE paths
$ProtoFiles = Get-ChildItem -Path $ProtoDir -Recurse -Filter *.proto | ForEach-Object {
    # Convert absolute -> relative path
    $relative = Resolve-Path -Relative $_.FullName
    $relative
}

foreach ($file in $ProtoFiles) {
    Write-Host "Generating Go code for $file..."

    protoc `
        -I $ProtoDir `
        --go_out=$OutDir `
        --go_opt=paths=source_relative `
        --go-grpc_out=$OutDir `
        --go-grpc_opt=paths=source_relative `
        $file
}

# Write-Host "✔️ Done. Generated files in '$OutDir'."
