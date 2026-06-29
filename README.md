# go-android-app

Exemplo mínimo de um aplicativo Android nativo escrito em **Go** usando
[`gomobile`](https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile).

O app abre uma janela nativa no Android, preenche a tela com uma cor
sólida (verde-azulado) e exibe um contador de FPS no canto. Não usa
Java/Kotlin — todo o código de UI é Go puro, compilado para um APK
diretamente executável.

## Estrutura

```
.
├── main.go                  # ponto de entrada do app (gomobile build)
├── go.mod
├── go.sum
└── .github/
    └── workflows/
        └── build2.yml       # CI: configura Go + Android NDK e gera o APK
```

## Build local (Linux/macOS)

Pré-requisitos:
- Go 1.21+
- Android SDK Platform 34 + Build-Tools 34.0.0
- Android NDK 26.x

```bash
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init
gomobile build -target=android/arm64 .
# gera ./go-android-app.apk
```

## Instalar no dispositivo

```bash
adb install -r go-android-app.apk
adb shell am start -n org.golang.app/.GoNativeActivity
```

## CI

O workflow `.github/workflows/build2.yml` pode ser disparado manualmente
(`workflow_dispatch`) ou a cada push em `main`. Ele:

1. Configura Go 1.21
2. Instala Android SDK + NDK
3. Instala e inicializa o `gomobile`
4. Executa `gomobile build -target=android/arm64,android/amd64`
5. Faz upload do APK gerado como artifact

## Licença

MIT
