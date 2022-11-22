
# vs_export
read visual studio 15/17/19/22 sln file,export clang compile_commands.json

```cmd
usage:vs_export options
           -s   path                        sln filename
           -c   configuration               project configuration,eg Debug|x64.
                                            default Debug|x64
```

## example

```cmd
vs_export.exe  -s NYWinHotspot.sln  -c "Debug|x64"
```

this can export a compile_commands.json. the compile_commands.json can used by clangd or ccls or some other cpp language server.
