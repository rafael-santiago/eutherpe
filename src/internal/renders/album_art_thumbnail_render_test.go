package renders

import (
    "internal/vars"
    "fmt"
    "testing"
)

func TestAlbumArtThumbnailRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedleAlbumArtThumbnail)
    output := AlbumArtThumbnailRender(templatedInput, eutherpeVars)
    if len(output) > 0 {
        t.Errorf("AlbumArtThumbnailRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.Player.NowPlaying.AlbumCover = "NoPassinNoPassinNoPassinDoQualquerCoisa(MasNaoDeUmaCoisaQualquer...)"
    output = AlbumArtThumbnailRender(templatedInput, eutherpeVars)
    if output != "<img src=\"data:image/umblauqualquerquenaovaiabrir;base64,Tm9QYXNzaW5Ob1Bhc3Npbk5vUGFzc2luRG9RdWFscXVlckNvaXNhKE1hc05hb0RlVW1hQ29pc2FRdWFscXVlci4uLik=\" width=50 height=50>" {
        t.Errorf("AlbumArtThumbnailRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.Player.NowPlaying.AlbumCover = "\x89PNG\r\n\x1A\n(PNG GOES HERE...)"
    output = AlbumArtThumbnailRender(templatedInput, eutherpeVars)
    if output != "<img src=\"data:image/png;base64,iVBORw0KGgooUE5HIEdPRVMgSEVSRS4uLik=\" width=50 height=50>" {
        t.Errorf("AlbumArtThumbnailRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.Player.NowPlaying.AlbumCover = "\xFF\xD8(JPEG GOES HERE...)\xFF\xD9"
    output = AlbumArtThumbnailRender(templatedInput, eutherpeVars)
    if output != "<img src=\"data:image/jpeg;base64,/9goSlBFRyBHT0VTIEhFUkUuLi4p/9k=\" width=50 height=50>" {
        t.Errorf("AlbumArtThumbnailRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.Player.NowPlaying.AlbumCover = "GIF87a(GIF GOES HERE...)"
    output = AlbumArtThumbnailRender(templatedInput, eutherpeVars)
    if output != "<img src=\"data:image/gif;base64,R0lGODdhKEdJRiBHT0VTIEhFUkUuLi4p\" width=50 height=50>" {
        t.Errorf("AlbumArtThumbnailRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.Player.NowPlaying.AlbumCover = "GIF89a(GIF GOES HERE...)"
    output = AlbumArtThumbnailRender(templatedInput, eutherpeVars)
    if output != "<img src=\"data:image/gif;base64,R0lGODlhKEdJRiBHT0VTIEhFUkUuLi4p\" width=50 height=50>" {
        t.Errorf("AlbumArtThumbnailRender() seems not to be rendering accordingly.\n")
    }
}
