//go:build !server

package opennox

import (
	"context"
	"errors"
	"fmt"
	"image"
	"math"
	"strings"

	"github.com/opennox/libs/client/seat"
	"github.com/opennox/libs/datapath"
	"github.com/opennox/libs/maps"

	"github.com/opennox/opennox/v1/client/gui"
	"github.com/opennox/opennox/v1/client/noxrender"
	noxflags "github.com/opennox/opennox/v1/common/flags"
	"github.com/opennox/opennox/v1/common/memmap"
	"github.com/opennox/opennox/v1/common/sound"
	"github.com/opennox/opennox/v1/legacy"
	"github.com/opennox/opennox/v1/legacy/common/alloc"
	"github.com/opennox/opennox/v1/server"
)

var (
	ticks815724 uint64
)

func (c *Client) drawAndPresent() {
	if nox_client_gui_flag_815132 != 0 {
		gui.AnimTick()
		noxflags.UnsetEngine(noxflags.EnginePause)
		c.generateMouseSparks()
	}
	if !noxflags.HasEngine(noxflags.EnginePause) {
		c.mainloopDrawAndPresent()
	}
}

func (c *Client) map_download_start() {
	c.nox_xxx_gameClearAll_467DF0(true)
	c.ShowMapDownload()
	c.mapsend.setDownloading(true)
	c.mapsend.setDownloadOK(true)
	if _, err := c.mapDownloadLoop(true); err != nil {
		c.mapsend.log.Error("download failed", "err", err)
	}
}

func (c *Client) mapDownloadLoop(first bool) (bool, error) {
	if !c.mapsend.Downloading() {
		if c.map_download_finish() == 0 {
			return true, errors.New("map download failed")
		}
		return true, nil
	}

	c.srv.SetRateLimit(30)
	c.processInput()
	sub_43CCA0()

	if first {
		ctx, cancel := context.WithCancel(context.Background())

		hport := server.InferHTTPPort(legacy.ClientGetServerPort())
		srv := fmt.Sprintf("%s:%d", clientGetServerHost(), hport)
		c.mapsend.log.Info("checking map download API on server", "addr", srv)
		cli, err := maps.NewClient(ctx, c.Log, srv)
		if err != nil {
			if err == maps.ErrAPIUnsupported {
				c.mapsend.log.Info("map API check", "err", err)
			} else {
				c.mapsend.log.Warn("cannot check map API", "err", err)
			}
			cancel()
			c.mapsend.native = true
		} else {
			c.mapsend.log.Info("map API supported")
			errc := make(chan error, 1)
			c.mapsend.cancel = cancel
			c.mapsend.done = errc
			c.mapsend.native = false

			name := strings.TrimSuffix(strings.ToLower(clientGetServerMap()), maps.Ext)
			c.mapsend.log.Info("download start", "proto", "http", "name", name)
			go func() {
				defer cli.Close()
				defer close(errc)

				err := cli.DownloadMap(ctx, datapath.Data(maps.Dir), name)
				if err != nil {
					c.mapsend.log.Error("download failed", "err", err)
					errc <- err
				} else {
					c.mapsend.log.Info("download complete")
				}
			}()
		}
	}

	if c.mapsend.native {
		if first {
			c.Nox_xxx_netRequestMap_43CA50()
		}
		if c.srv.Frame()%30 != 0 { // TODO: shouldn't it be == 0?
			c.Nox_xxx_netKeepAliveSocket_43CA20()
		}
	} else {
		select {
		case err := <-c.mapsend.done:
			c.mapsend.setDownloading(false)
			if err != nil {
				c.mapsend.setDownloadOK(false)
				return true, err
			}
			c.mapsend.setDownloadOK(true)
			c.mapsend.setProgress(100)
			if c.map_download_finish() == 0 {
				return true, errors.New("map download failed")
			}
			return true, nil
		default:
		}
	}
	if !c.mapsend.Downloading() {
		if c.map_download_finish() == 0 {
			return true, errors.New("map download failed")
		}
		return true, nil
	}

	c.GUI.Draw()
	c.nox_client_drawCursorAndTooltips_477830()
	c.copyPixBuffer()

	if !c.mapsend.Downloading() {
		if c.map_download_finish() == 0 {
			return true, errors.New("map download failed")
		}
		return true, nil
	}
	// continue
	return false, nil
}

func (c *Client) mainloopDrawAndPresent() {
	sub_437180()
	if legacy.Get_nox_client_gui_flag_1556112() == 0 {
		c.GUI.Draw() // Draw game windows
	}
	c.DrawSparks()
	if !noxflags.HasEngine(noxflags.EngineNoRendering) || noxflags.HasEngine(noxflags.EngineFlag9) || nox_client_gui_flag_815132 != 0 {
		c.nox_client_drawCursorAndTooltips_477830() // Draw cursor
	}
	c.r.DrawFade(true)
	c.maybeScreenshot()
	if !noxflags.HasEngine(noxflags.EngineNoRendering) || noxflags.HasEngine(noxflags.EngineFlag9) || nox_client_gui_flag_815132 != 0 {
		c.copyPixBuffer()
	}
}

func (c *Client) DrawSparks() {
	if nox_client_gui_flag_815132 != 0 {
		sz := videoGetWindowSize()
		vp, rdrFree := alloc.New(noxrender.Viewport{})
		defer rdrFree()
		vp.Screen.Min = image.Pt(0, 0)
		vp.Screen.Max = sz
		vp.Size = sz
		legacy.Nox_client_screenParticlesDraw_431720(vp)
	} else {
		vp := c.Viewport()
		legacy.Nox_client_screenParticlesDraw_431720(vp)
	}
}

func (c *Client) generateMouseSparks() {
	if memmap.Uint32(0x5D4594, 816408) != 0 {
		return
	}

	mpos := c.Inp.GetMousePos()
	// emit sparks when passing a certain distance
	const distanceSparks = 0.25
	dx := mpos.X - int(memmap.Uint32(0x5D4594, 816420))
	dy := mpos.Y - int(memmap.Uint32(0x5D4594, 816424))
	r2 := dx*dx + dy*dy
	if memmap.Uint32(0x5D4594, 816428) != 0 {
		cnt := int(math.Sqrt(float64(r2)) * distanceSparks)
		for i := cnt; i > 0; i-- {
			v6 := c.srv.Rand.Other.Int(0, 100)
			v7 := int(memmap.Uint32(0x5D4594, 816420)) + dx*v6/100
			v9 := int(memmap.Uint32(0x5D4594, 816424)) + dy*v6/100
			v23 := c.srv.Rand.Other.Int(2, 5)
			v22 := c.srv.Rand.Other.Int(2, 5)
			v21 := c.srv.Rand.Other.Int(-7, 2)
			v10 := c.srv.Rand.Other.Int(-5, 5)
			legacy.Nox_client_newScreenParticle_431540(4, v7, v9, v10, v21, 1, v22, v23, 2, 1)
		}
		if r2 < 10 {
			*memmap.PtrUint32(0x5D4594, 816428) = 0
		}
		*memmap.PtrUint32(0x5D4594, 816420) = uint32(mpos.X)
		*memmap.PtrUint32(0x5D4594, 816424) = uint32(mpos.Y)
	} else if r2 > 64 {
		*memmap.PtrUint32(0x5D4594, 816428) = 1
	}
	// explode with sparks when clicking
	const explosionSparks = 75
	if c.Inp.IsMousePressed(seat.MouseButtonLeft) {
		c.srv.Rand.Other.Int(0, 2)
		if memmap.Uint32(0x5D4594, 816416) == 0 {
			*memmap.PtrUint32(0x5D4594, 816416) = 1
			clientPlaySoundSpecial(sound.SoundShellMouseBoom, 100)
			for i := explosionSparks; i > 0; i-- {
				v12 := c.srv.Rand.Other.Int(0, 255)
				v13 := c.srv.Rand.Other.Int(6, 12)
				pos := sincosTable16[v12].Mul(v13).Div(16).Add(image.Point{Y: -6})
				v24 := c.srv.Rand.Other.Int(2, 5)
				v16 := c.srv.Rand.Other.Int(2, 5)
				legacy.Nox_client_newScreenParticle_431540(4, pos.X+mpos.X, pos.Y+mpos.Y, pos.X, pos.Y, 1, v16, v24, 2, 1)
			}
		}
	} else {
		*memmap.PtrUint32(0x5D4594, 816416) = 0
	}
}
