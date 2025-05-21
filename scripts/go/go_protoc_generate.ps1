$PROTO_DIR = "proto"
$OUT_DIR = "proto/generated/go"

$protoFiles = Get-ChildItem -Path $PROTO_DIR -Recurse -Filter *.proto

foreach ($protoFile in $protoFiles) {
    # Получаем относительный путь
    $relativePath = $protoFile.FullName.Substring((Resolve-Path $PROTO_DIR).Path.Length + 1)
    $relativePath = $relativePath -replace "\\", "/"

    # Извлекаем имя поддиректории
    $packageName = $relativePath.Split(".")[0] + "v1"
    $outPath = Join-Path $OUT_DIR $packageName\

    if (!(Test-Path $outPath)) {
        New-Item -ItemType Directory -Path $outPath | Out-Null
    }

    protoc -I $PROTO_DIR `
        "$PROTO_DIR/$relativePath" `
        --go_out="$outPath" `
        --go_opt=paths=source_relative `
        --go-grpc_out="$outPath" `
        --go-grpc_opt=paths=source_relative
}
