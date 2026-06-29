// Exemplo de app Android escrito em Go usando golang.org/x/mobile.
//
// Este app abre uma janela nativa no Android, limpa a tela com uma cor
// de fundo (verde-azulado) e imprime eventos de lifecycle no logcat.
// Compilado com `gomobile build` gera um APK diretamente executável
// no dispositivo (não é uma biblioteca de bindings).
//
// Requer:
//   - Go >= 1.21
//   - Android SDK + NDK
//   - gomobile + gobind
package main

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

func main() {
	app.Main(func(a app.App) {
		var (
			glctx   gl.Context
			sz      size.Event
			fps     *debug.FPS
			images  *glutil.Images
		)

		for {
			select {
			case e := <-a.Events():
				switch e := a.Filter(e).(type) {
				case lifecycle.Event:
					switch e.Crosses(lifecycle.StageVisible) {
					case lifecycle.CrossOn:
						glctx, _ = e.DrawContext.(gl.Context)
						images = glutil.NewImages(glctx)
						fps = debug.NewFPS(images)
					case lifecycle.CrossOff:
						if images != nil {
							images.Release()
							images = nil
						}
						glctx = nil
					}
				case size.Event:
					sz = e
				case paint.Event:
					if glctx == nil || sz.WidthPx == 0 {
						continue
					}
					// Cor de fundo: verde-azulado (R=0.10, G=0.40, B=0.55)
					glctx.ClearColor(0.10, 0.40, 0.55, 1.0)
					glctx.Clear(gl.COLOR_BUFFER_BIT)
					if fps != nil {
						fps.Draw(sz)
					}
					a.Publish()
					a.Send(paint.Event{})
				}
			}
		}
	})
}
